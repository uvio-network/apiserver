package votereconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
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
				return nil, tracer.Maskf(StakeLifecyclePendingError, "%s", inp[i].Owner)
			}

			// Votes of kind "stake" must comply with the lifecycle of their
			// referenced claim object, e.g. you cannot stake on a resolve.
			if inp[i].Kind == "stake" && !cla.Lifecycle.Is(objectlabel.LifecycleDispute, objectlabel.LifecyclePropose) {
				return nil, tracer.Maskf(StakeLifecycleInvalidError, "%s", cla.Lifecycle.Data)
			}

			// Votes of kind "truth" must comply with the lifecycle of their
			// referenced claim object, e.g. you can only vote on a resolve.
			if inp[i].Kind == "truth" && !cla.Lifecycle.Is(objectlabel.LifecycleResolve) {
				return nil, tracer.Maskf(TruthLifecycleInvalidError, "%s", cla.Lifecycle.Data)
			}

			// Ensure no votes can be cast anymore on claims that have already expired.
			if now.After(cla.Expiry) {
				return nil, tracer.Maskf(ClaimAlreadyExpiredError, "%d", cla.Expiry.Unix())
			}
		}

		// Prevent multiple votes by the same user on the same claim. For market
		// resolutions it stands one user one vote. This is also enforced in the
		// underlying smart contracts.
		if inp[i].Kind == "truth" {
			var cou int64
			{
				cou, err = r.red.Sorted().Metric().Count(votOwnCla(inp[i].Owner, inp[i].Claim))
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			if cou != 0 {
				return nil, tracer.Maskf(VoteAlreadyExistsError, "%s", cla.ID)
			}
		}
	}

	return inp, nil
}
