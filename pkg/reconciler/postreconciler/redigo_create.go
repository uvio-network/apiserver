package postreconciler

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/labelname"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/uvio-network/apiserver/pkg/summary"
	"github.com/xh3b4sd/tracer"
)

const (
	oneWeek = time.Hour * 24 * 7
)

func (r *Redigo) CreateBalance(res *poststorage.Object) (*poststorage.Object, error) {
	var err error

	var bal *poststorage.Object
	{
		bal = &poststorage.Object{
			Chain:    res.Chain,
			Contract: res.Contract,
			Kind:     "claim",
			Labels:   res.Labels,
			Lifecycle: objectfield.Lifecycle{
				Data: objectlabel.LifecycleBalance,
			},
			Owner:  objectid.System(),
			Parent: res.ID,
			Text:   "# Market Settlement\n\nThe process of updating user balances has begun and is waiting for onchain confirmation.",
		}
	}

	var out []*poststorage.Object
	{
		out, err = r.CreatePost([]*poststorage.Object{bal})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = r.sto.Post().CreatePost(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out[0], nil
}

func (r *Redigo) CreatePost(inp []*poststorage.Object) ([]*poststorage.Object, error) {
	var err error

	for i := range inp {
		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			if inp[i].Kind == "claim" {
				inp[i].Lifecycle.Time = now
				inp[i].Summary = []float64{0, 0, 0, 0}
			}

			if inp[i].Kind == "comment" {
				inp[i].Summary = []float64{0, 0}
			}
		}

		{
			err := inp[i].Verify()
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			inp[i].Created = now
			inp[i].ID = objectid.Random(objectid.Time(now))
			inp[i].Labels = generic.Func(inp[i].Labels, labelname.Format)
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
		var par *poststorage.Object
		if inp[i].Kind == "claim" && inp[i].Lifecycle.Is(objectlabel.LifecyclePropose) {
			inp[i].Tree = objectid.Random(objectid.Time(now))
		} else {
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

		// Lookup whether the post owner was selected by the truth sampling process.
		// If the post being created is for instance a comment, and if that comment
		// is referencing a resolve, then "sel" will be set to true, which prevents
		// the MarketParticipationError to be returned below.
		var sel bool
		if par != nil {
			for _, v := range par.Samples {
				if v == string(inp[i].Owner) {
					sel = true
					break
				}
			}
		}

		// When creating posts of kind "claim", we want to ensure that claims cannot
		// be created with expiries that point to the past. Claim expiries are only
		// relevant for claims of lifecycle phase "dispute", "propose" and
		// "resolve".
		if inp[i].Kind == "claim" && inp[i].Lifecycle.Is(objectlabel.LifecycleDispute, objectlabel.LifecyclePropose, objectlabel.LifecycleResolve) {
			if inp[i].Expiry.Before(now) {
				return nil, tracer.Maskf(ClaimExpiryPastError, "%s = %d", inp[i].ID, inp[i].Expiry.Unix())
			}
		}

		if inp[i].Kind == "claim" && inp[i].Lifecycle.Is(objectlabel.LifecycleDispute) {
			// Ensure that all involved claims are facilitated by the same smart
			// contract on the same blockchain network.
			if inp[i].Chain != par.Chain || inp[i].Contract != par.Contract {
				return nil, tracer.Maskf(DisputeContractError, "%s / %s", inp[i].ID, par.ID)
			}

			// Verify whether the resolve can still be disputed, which means the
			// resolve must still be within its challenge window of 7 standard days
			// from its expiry.
			if now.After(par.Expiry.Add(oneWeek)) {
				return nil, tracer.Maskf(DisputeChallengeError, "%s", par.ID)
			}

			// Ensure that disputes can only be created by referencing resolves.
			if !par.Lifecycle.Is(objectlabel.LifecycleResolve) {
				return nil, tracer.Maskf(DisputeLifecycleError, "%s", par.ID)
			}

			// Ensure disputes can only be created on resolves with valid resolution.
			// If resolves had no votes, or an equal amount of votes on each side of
			// the market, then those resolutions are considered invalid. Those
			// invalid resolutions are considered definitive and binding, and can
			// therefore not be disputed.
			if !summary.Verify(par) {
				return nil, tracer.Maskf(DisputeResolutionError, "%s", par.ID)
			}

			var tre poststorage.Slicer
			{
				tre, err = r.sto.Post().SearchTree([]objectid.ID{par.Tree})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			var dis []*poststorage.Object
			var pro []*poststorage.Object
			var res []*poststorage.Object
			{
				dis = tre.ObjectLifecycle(objectlabel.LifecycleDispute)
				pro = tre.ObjectLifecycle(objectlabel.LifecyclePropose)
				res = tre.ObjectLifecycle(objectlabel.LifecycleResolve)
			}

			if len(pro) != 1 {
				return nil, tracer.Mask(runtime.ExecutionFailedError)
			}

			// Enforce the max dispute limit for our offchain resources. The same
			// limitation is enforced in the smart contracts.
			if len(dis) >= 2 {
				return nil, tracer.Maskf(DisputeLimitError, "%s / %s", pro[0].ID, par.ID)
			}

			// Ensure the dispute that we are creating has the same token definition
			// as the original propose.
			if inp[i].Token != pro[0].Token {
				return nil, tracer.Maskf(DisputeContractError, "%s / %s", inp[i].ID, par.ID)
			}

			// Below we push back the resolve expiries within this tree, in order to
			// prevent our queued background jobs to prematurely reconcile system
			// state.
			{
				err = r.sto.Post().DeleteExpiry(res)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			// Since this here is the reconciliation of dispute creation, we will
			// either find one or two resolves, both of which should have the same
			// queued expiry. And since we are just now creating a dispute, those
			// existing resolves must be deferred again. Their new queued expiries
			// should be aligned with the expiry of the future resolve, which is not
			// yet created, but which will be created once this dispute that we create
			// here expired on its own. And since resolves have a voting period of 7
			// standard days, and since resolves have a challenge window of 7 standard
			// days, the existing resolves must be pushed back in the expiry queue by
			// 2 standard weeks.
			for j := range res {
				res[j].Expiry = inp[i].Expiry.Add(2 * oneWeek)
			}

			// Note that pushing back the queued expiry does not modify the
			// Post.Expiry field in the resolve objects. All we want to do here is to
			// defer our background job execution.
			{
				err = r.sto.Post().CreateExpiry(res)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}

		// Ensure that the calling user has in fact skin in the game. If we cannot
		// find a single vote object referencing the provided parent claim for the
		// calling user, and if the calling user has not been selected by the truth
		// sampling process, then we return an error, because only users who have
		// actually participated in any given market are allowed to express their
		// opinions about those markets.
		if inp[i].Kind == "comment" {
			var vot []*votestorage.Object
			{
				vot, err = r.sto.Vote().SearchOwnerClaim([]objectid.ID{inp[i].Owner}, []objectid.ID{inp[i].Parent})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			if !sel && len(vot) == 0 {
				return nil, tracer.Mask(MarketParticipationError)
			}

			// Update the vote summary of the comment according to all existing votes.
			for _, x := range vot {
				inp[i] = summary.Update(inp[i], x)
			}
		}
	}

	return inp, nil
}

func (r *Redigo) CreateResolve(pod *poststorage.Object, exp time.Time) (*poststorage.Object, error) {
	var err error

	var res *poststorage.Object
	{
		res = &poststorage.Object{
			Chain:    pod.Chain,
			Contract: pod.Contract,
			Expiry:   exp,
			Kind:     "claim",
			Labels:   pod.Labels,
			Lifecycle: objectfield.Lifecycle{
				Data: objectlabel.LifecycleResolve,
			},
			Owner:  objectid.System(),
			Parent: pod.ID,
			Text:   "# Market Resolution\n\nThe random truth sampling process has begun and is waiting for onchain confirmation.",
		}
	}

	var out []*poststorage.Object
	{
		out, err = r.CreatePost([]*poststorage.Object{res})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = r.sto.Post().CreatePost(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out[0], nil
}

func (r *Redigo) verifyParent(par objectid.ID) (*poststorage.Object, error) {
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
		return nil, tracer.Maskf(ParentKindError, "%s = %s", obj[0].ID, obj[0].Kind)
	}

	// Here we ensure that posts cannot be created on any pending resources.
	if obj[0].Lifecycle.Pending() {
		return nil, tracer.Maskf(ParentPendingError, "%s = %s", obj[0].ID, obj[0].Lifecycle)
	}

	return obj[0], nil
}
