package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectstatus"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *vote.UpdateI) (*vote.UpdateO, error) {
	var err error

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
	// Search for all relevant vote objects.
	//

	var vot []*votestorage.Object
	{
		vot, err = h.sto.Vote().SearchVote(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Verify the ownership of the given resources.
	//

	for i := range vot {
		if vot[i].Owner != use {
			return nil, tracer.Maskf(runtime.UserNotOwnerError, "%s != %s", vot[i].Owner, use)
		}
	}

	//
	// Update the vote hash.
	//

	if len(has) != 0 {
		if len(ids) != len(has) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(ids), len(has))
		}

		{
			vot, err = h.rec.Vote().UpdateHash(vot, has)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	//
	// Update the given resources.
	//

	{
		err = h.sto.Vote().UpdateVote(vot)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *vote.UpdateO
	{
		res = &vote.UpdateO{}
	}

	for range vot {
		res.Object = append(res.Object, &vote.UpdateO_Object{
			Intern: &vote.UpdateO_Object_Intern{
				Status: objectstatus.Updated,
			},
		})
	}

	return res, nil
}
