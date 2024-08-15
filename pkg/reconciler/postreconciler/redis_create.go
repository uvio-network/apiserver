package postreconciler

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/labelname"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreatePost(inp []*poststorage.Object) ([]*poststorage.Object, error) {
	var err error

	for i := range inp {
		{
			if inp[i].Kind == "claim" {
				inp[i].Votes = []float64{0, 0, 0, 0}
			}

			if inp[i].Kind == "comment" {
				inp[i].Votes = []float64{0, 0}
			}
		}

		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Created = now
			inp[i].ID = objectid.Random(objectid.Time(now))
			inp[i].Labels = generic.Func(inp[i].Labels, labelname.Format)
			inp[i].Lifecycle.Time = now
			inp[i].Text = strings.TrimSpace(inp[i].Text)
			inp[i].Token = strings.TrimSpace(inp[i].Token)
		}

		// Lookup the existing tree ID, or create a new one. Any proposed claim
		// effectively starts a new tree. So for every post of kind "claim" that has
		// lifecycle "propose" we generate a new tree ID. For all other posts,
		// whether they are of kind "claim" or kind "comment", we fetch the post
		// object provided in the parent field, and take the existing tree ID from
		// there. Note that we are indirectly validating here that the given parent
		// ID does in fact exist.
		if inp[i].Kind == "claim" && inp[i].Lifecycle.Is(objectlabel.LifecyclePropose) {
			inp[i].Tree = objectid.Random(objectid.Time(now))
		} else {
			var par *poststorage.Object
			{
				par, err = r.verifyParent(inp[i].Parent)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			{
				inp[i].Tree = par.Tree
			}
		}

		// Ensure that the calling user has in fact skin in the game. If we cannot
		// find a single vote object referencing the provided parent claim for the
		// calling user, then we return an error, because only users who have
		// actually participated in any given market are allowed to express their
		// opinions about it.
		if inp[i].Kind == "comment" {
			var vot []*votestorage.Object
			{
				vot, err = r.sto.Vote().SearchOwnerClaim([]objectid.ID{inp[i].Owner}, []objectid.ID{inp[i].Parent})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			if len(vot) == 0 {
				return nil, tracer.Mask(PostOwnerVoteError)
			}

			// Update the vote summary of the comment according to all existing votes.
			for _, x := range vot {
				inp[i] = updateVotes(inp[i], x)
			}
		}
	}

	return inp, nil
}

func (r *Redis) verifyParent(par objectid.ID) (*poststorage.Object, error) {
	var err error

	var obj []*poststorage.Object
	{
		obj, err = r.sto.Post().SearchPost([]objectid.ID{par})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(obj) != 1 {
		return nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	// Here we ensure that comments cannot comment on comments.
	if obj[0].Kind != "claim" {
		return nil, tracer.Maskf(PostParentKindError, "kind=%s", obj[0].Kind)
	}

	return obj[0], nil
}
