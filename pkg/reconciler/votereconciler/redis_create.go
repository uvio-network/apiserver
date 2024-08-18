package votereconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateVote(inp []*votestorage.Object) ([]*votestorage.Object, error) {
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
			cla, err = r.verifyClaim(inp[i].Claim)
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
			// referenced claim object.
			if inp[i].Kind == "stake" && !cla.Lifecycle.Is(objectlabel.LifecycleAdjourn, objectlabel.LifecycleDispute, objectlabel.LifecycleNullify, objectlabel.LifecyclePropose) {
				return nil, tracer.Maskf(StakeLifecycleInvalidError, cla.Lifecycle.String())
			}

			// Votes of kind "truth" must comply with the lifecycle of their
			// referenced claim object.
			if inp[i].Kind == "truth" && !cla.Lifecycle.Is(objectlabel.LifecycleResolve) {
				return nil, tracer.Mask(TruthLifecycleInvalidError)
			}
		}

		// Ensure no votes can be cast anymore on claims that have already expired.
		{
			if now.After(cla.Expiry) {
				return nil, tracer.Maskf(ClaimAlreadyExpiredError, "%d", cla.Expiry.Unix())
			}
		}

		// Once all cross validation is done we can proceed with cross mutation. One
		// important thing we need to do for all users is to update their staked
		// token balances in their respective user objects.
		{
			var use []*userstorage.Object
			{
				use, err = r.sto.User().SearchUser([]objectid.ID{inp[i].Owner})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			for j := range use {
				use[j].Staked.Data[cla.Token] += inp[i].Value
				use[j].Staked.Time = now
			}

			{
				err = r.sto.User().UpdateUser(use)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}
	}

	return inp, nil
}

func (r *Redis) verifyClaim(cla objectid.ID) (*poststorage.Object, error) {
	var err error

	var obj []*poststorage.Object
	{
		obj, err = r.sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Every vote object must reference an existing claim object.
	if len(obj) != 1 {
		return nil, tracer.Maskf(VoteClaimNotFoundError, cla.String())
	}

	return obj[0], nil
}
