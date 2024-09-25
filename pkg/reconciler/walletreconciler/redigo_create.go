package walletreconciler

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) CreateWallet(inp []*walletstorage.Object) ([]*walletstorage.Object, error) {
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
			inp[i].Address = common.HexToAddress(inp[i].Address).Hex() // ensure consistent string format
			inp[i].Created = now
			inp[i].ID = objectid.Random(objectid.Time(now))
		}
	}

	return inp, nil
}
