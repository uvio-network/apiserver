package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectstatus"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *vote.UpdateI) (*vote.UpdateO, error) {
	var err error
	var out []string

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var ids []objectid.ID
	var has []string
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" && x.Public != nil && x.Public.Hash != "" {
			ids = append(ids, objectid.ID(x.Intern.Id))
			has = append(has, x.Public.Hash)
		}
	}

	//
	// Update the vote hash.
	//

	if len(has) != 0 {
		var upd []*votestorage.Object
		{
			upd, err = h.rec.Vote().UpdateHash(use, ids, has)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			err = h.sto.Vote().UpdateVote(upd)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, objectstatus.Updated)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *vote.UpdateO
	{
		res = &vote.UpdateO{}
	}

	for _, x := range out {
		res.Object = append(res.Object, &vote.UpdateO_Object{
			Intern: &vote.UpdateO_Object_Intern{
				Status: x,
			},
		})
	}

	return res, nil
}
