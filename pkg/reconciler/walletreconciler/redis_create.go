package walletreconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) CreateWallet(inp []*walletstorage.Object) ([]*walletstorage.Object, error) {
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

		if inp[i].Active {
			var sli walletstorage.Slicer
			{
				sli, err = r.sto.Wallet().SearchOwner([]objectid.ID{inp[i].Owner})
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			var upd []*walletstorage.Object
			{
				upd = sli.ObjectActive(true).ObjectKind(inp[i].Kind)
			}

			for j := range upd {
				upd[j].Active = false
			}

			if len(upd) != 0 {
				err = r.sto.Wallet().UpdateWallet(upd)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}
		}
	}

	return inp, nil
}
