package votereconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) CreateVote(inp []*votestorage.Object) ([]*votestorage.Object, error) {
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
			inp[i].Lifecycle.Time = now
		}

		var cla *poststorage.Object
		{
			cla, err = r.searchClaim(inp[i].Claim)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			// Votes can generally not be cast on claims that are still stuck in
			// lifecycle phase "pending". The exception is the proposer of the
			// referenced claim itself. As long as a new claim is still "pending", the
			// only user who can vote on the proposed claim is the proposer itself.
			// Here we make sure that users can propose claims on the platform without
			// having the onchain confirmation right away.
			if inp[i].Kind == "stake" && cla.Lifecycle.Pending() && inp[i].Owner != cla.Owner {
				return nil, tracer.Maskf(ClaimLifecyclePendingError, "%s != %s", inp[i].Owner, cla.Owner)
			}

			// Votes of kind "stake" must comply with the lifecycle of their
			// referenced claim object, e.g. you cannot stake on a resolve.
			if inp[i].Kind == "stake" && !cla.Lifecycle.Is(objectlabel.LifecycleDispute, objectlabel.LifecyclePropose) {
				return nil, tracer.Maskf(ClaimLifecycleInvalidError, "%s = %s", cla.ID, cla.Lifecycle.Data)
			}

			// Votes of kind "truth" must comply with the lifecycle of their
			// referenced claim object, e.g. you can only vote on a resolve.
			if inp[i].Kind == "truth" && !cla.Lifecycle.Is(objectlabel.LifecycleResolve) {
				return nil, tracer.Maskf(TruthLifecycleInvalidError, "%s = %s", cla.ID, cla.Lifecycle.Data)
			}

			// Ensure no votes can be cast anymore on claims that have already expired.
			if now.After(cla.Expiry) {
				return nil, tracer.Maskf(ClaimAlreadyExpiredError, "%s = %d", cla.ID, cla.Expiry.Unix())
			}
		}

		var vot *votestorage.Object
		{
			vot, err = r.searchVote(inp[i].Owner, inp[i].Claim)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Prevent the user from staking if the latest stake is still pending.
		// Before the user can stake more reputation, their previous stake must be
		// reconciled first.
		if inp[i].Kind == "stake" && vot != nil {
			if vot.Lifecycle.Pending() {
				return nil, tracer.Maskf(StakeLifecyclePendingError, "%s", cla.ID)
			}
		}

		// Prevent multiple votes by the same user on the same claim from being
		// created. For market resolutions it stands "one user one vote". This is
		// also enforced in the underlying smart contracts.
		if inp[i].Kind == "truth" && vot != nil {
			if vot.Lifecycle.Pending() {
				return nil, tracer.Maskf(VoteLifecyclePendingError, "%s", cla.ID)
			}
			if !vot.Lifecycle.Pending() {
				return nil, tracer.Maskf(VoteAlreadyExistsError, "%s", cla.ID)
			}
		}
	}

	return inp, nil
}

func (r *Redigo) searchVote(use objectid.ID, cla objectid.ID) (*votestorage.Object, error) {
	var err error

	var val []string
	{
		val, err = r.red.Sorted().Search().Order(votOwnCla(use, cla), -1, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(val) == 0 {
		return nil, nil
	}

	var out []*votestorage.Object
	{
		lis, err := r.sto.Vote().SearchVote(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	if len(out) == 0 {
		return nil, nil
	}

	if len(out) == 1 {
		return out[0], nil
	}

	return nil, tracer.Mask(runtime.ExecutionFailedError)
}
