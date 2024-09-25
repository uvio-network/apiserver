package postreconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) DeletePost(inp []*poststorage.Object) ([]*poststorage.Object, error) {
	for i := range inp {
		if inp[i].Kind == "claim" {
			var now time.Time
			{
				now = time.Now().UTC()
			}

			// Claims can only be deleted as long as they have not been confirmed
			// onchain.
			if !inp[i].Lifecycle.Pending() {
				return nil, tracer.Maskf(ClaimLifecycleOnchainError, inp[i].ID.String())
			}

			// Ensure no claims can be removed anymore that have already expired.
			if now.After(inp[i].Expiry) {
				return nil, tracer.Maskf(ClaimAlreadyExpiredError, "%d", inp[i].Expiry.Unix())
			}
		}
	}

	return inp, nil
}
