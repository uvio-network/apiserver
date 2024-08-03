package votevalidator

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) Create(inp []*votestorage.Object) ([]*votestorage.Object, error) {
	var err error

	for i := range inp {
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
		}

		{
			var obj []*poststorage.Object
			{
				obj, err = r.sto.Post().SearchPost([]objectid.ID{inp[i].Claim})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Every vote object must reference an existing claim object.
			if len(obj) != 1 {
				return nil, tracer.Maskf(VoteClaimNotFoundError, inp[i].Claim.String())
			}

			// Votes of kind "stake" must comply with the lifecycle of their
			// referenced claim object.
			if inp[i].Kind == "stake" {
				if obj[0].Lifecycle != "adjourn" && obj[0].Lifecycle != "dispute" && obj[0].Lifecycle != "nullify" && obj[0].Lifecycle != "propose" {
					return nil, tracer.Mask(StakeLifecycleInvalidError)
				}
			}

			// Votes of kind "truth" must comply with the lifecycle of their
			// referenced claim object.
			if inp[i].Kind == "truth" {
				if obj[0].Lifecycle != "resolve" {
					return nil, tracer.Mask(TruthLifecycleInvalidError)
				}
			}
		}
	}

	return inp, nil
}
