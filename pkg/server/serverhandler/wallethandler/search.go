package wallethandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *wallet.SearchI) (*wallet.SearchO, error) {
	var out []*walletstorage.Object

	var own []objectid.ID
	var wal []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" {
			wal = append(wal, objectid.ID(x.Intern.Id))
		}
		if x.Intern != nil && x.Intern.Owner != "" {
			if x.Intern.Owner == "self" {
				own = append(own, userid.FromContext(ctx))
			} else {
				own = append(own, objectid.ID(x.Intern.Owner))
			}
		}
	}

	//
	// Search wallets by user ID.
	//

	if len(own) != 0 {
		lis, err := h.sto.Wallet().SearchOwner(own)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search wallets by wallet ID.
	//

	if len(wal) != 0 {
		lis, err := h.sto.Wallet().SearchWallet(wal)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct the RPC response.
	//

	var res *wallet.SearchO
	{
		res = &wallet.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "wallet",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		res.Object = append(res.Object, &wallet.SearchO_Object{
			Intern: &wallet.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
				Owner:   x.Owner.String(),
			},
			Public: &wallet.SearchO_Object_Public{
				Active:      converter.BoolToString(x.Active),
				Address:     x.Address,
				Description: x.Description,
				Kind:        x.Kind,
				Provider:    x.Provider,
			},
		})
	}

	return res, nil
}
