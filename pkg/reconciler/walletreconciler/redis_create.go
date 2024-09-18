package walletreconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateWallet(inp []*walletstorage.Object) ([]*walletstorage.Object, error) {
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
	}

	return inp, nil
}
