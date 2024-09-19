package votereconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) DeleteVote(inp []*votestorage.Object) ([]*votestorage.Object, error) {
	var err error

	for i := range inp {
		var now time.Time
		{
			now = time.Now().UTC()
		}

		var cla *poststorage.Object
		{
			cla, err = r.searchClaim(inp[i].Claim)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Votes can only be deleted as long as they have not been confirmed
		// onchain.
		if !cla.Lifecycle.Pending() {
			return nil, tracer.Maskf(VoteLifecycleOnchainError, inp[i].ID.String())
		}

		// Ensure no votes can be removed anymore on claims that have already expired.
		if now.After(cla.Expiry) {
			return nil, tracer.Maskf(ClaimAlreadyExpiredError, "%d", cla.Expiry.Unix())
		}
	}

	return inp, nil
}
