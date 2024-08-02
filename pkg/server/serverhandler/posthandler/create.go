package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *post.CreateI) (*post.CreateO, error) {
	var err error

	var inp []*poststorage.Object
	for _, x := range req.Object {
		inp = append(inp, &poststorage.Object{
			Expiry:    converter.StringToTime(x.Public.Expiry),
			Kind:      x.Public.Kind,
			Labels:    converter.StringToSlice(x.Public.Labels),
			Lifecycle: x.Public.Lifecycle,
			Option:    converter.StringToBool(x.Public.Option),
			Owner:     userid.FromContext(ctx),
			Parent:    objectid.ID(x.Public.Parent),
			Stake:     converter.StringToFloat(x.Public.Stake),
			Text:      x.Public.Text,
			Token:     x.Public.Token,
		})
	}

	//
	// Create the given resources.
	//

	{
		inp, err = h.val.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*poststorage.Object
	{
		out, err = h.sto.Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *post.CreateO
	{
		res = &post.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &post.CreateO_Object{
			Intern: &post.CreateO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
			},
		})
	}

	return res, nil
}
