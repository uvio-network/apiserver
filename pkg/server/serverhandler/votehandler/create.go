package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *vote.CreateI) (*vote.CreateO, error) {
	var err error

	var inp []*votestorage.Object
	for _, x := range req.Object {
		inp = append(inp, &votestorage.Object{
			Chain: x.Public.Chain,
			Claim: objectid.ID(x.Public.Claim),
			Kind:  x.Public.Kind,
			Lifecycle: objectfield.Lifecycle{
				Data: objectlabel.DesiredLifecycle(x.Public.Lifecycle),
				Hash: converter.StringToSlice(x.Public.Hash),
			},
			Meta:   x.Public.Meta,
			Option: converter.StringToBool(x.Public.Option),
			Owner:  userid.FromContext(ctx),
			Value:  converter.StringToFloat(x.Public.Value),
		})
	}

	//
	// Create the given resources.
	//

	var out []*votestorage.Object
	{
		out, err = h.rec.Vote().CreateVote(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = h.sto.Vote().CreateVote(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// update the vote summary for the referenced post
	//

	var pos []*poststorage.Object
	{
		pos, err = h.rec.Post().UpdateVotes(out)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		err = h.sto.Post().UpdatePost(pos)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

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
