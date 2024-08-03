package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *vote.CreateI) (*vote.CreateO, error) {
	var err error

	var inp []*votestorage.Object
	for _, x := range req.Object {
		inp = append(inp, &votestorage.Object{
			Claim:  objectid.ID(x.Public.Claim),
			Kind:   x.Public.Kind,
			Option: converter.StringToBool(x.Public.Option),
			Owner:  userid.FromContext(ctx),
			Value:  converter.StringToFloat(x.Public.Value),
		})
	}

	//
	// Create the given resources.
	//

	{
		inp, err = h.val.Vote().Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*votestorage.Object
	{
		out, err = h.sto.Vote().Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// TODO update vote summary for claims and comments

	//
	// Construct the RPC response.
	//

	var res *vote.CreateO
	{
		res = &vote.CreateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &vote.CreateO_Object{
			Intern: &vote.CreateO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
			},
		})
	}

	return res, nil
}
