package wallethandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *wallet.CreateI) (*wallet.CreateO, error) {
	var err error

	var inp []*walletstorage.Object
	for _, x := range req.Object {
		inp = append(inp, &walletstorage.Object{
			Active:      converter.StringToBool(x.Public.Active),
			Address:     x.Public.Address,
			Description: x.Public.Description,
			Kind:        x.Public.Kind,
			Owner:       userid.FromContext(ctx),
		})
	}

	//
	// Create the given resources.
	//

	var out []*walletstorage.Object
	{
		out, err = h.rec.Wallet().CreateWallet(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = h.sto.Wallet().CreateWallet(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *wallet.CreateO
	{
		res = &wallet.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &wallet.CreateO_Object{
			Intern: &wallet.CreateO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
			},
		})
	}

	return res, nil
}
