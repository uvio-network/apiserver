package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectstatus"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *vote.UpdateI) (*vote.UpdateO, error) {
	var err error

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var ids []objectid.ID
	var hsh []string
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" && x.Public != nil && x.Public.Hash != "" {
			ids = append(ids, objectid.ID(x.Intern.Id))
			hsh = append(hsh, x.Public.Hash)
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

	if len(hsh) != 0 {
		if len(ids) != len(hsh) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(ids), len(hsh))
		}

		{
			vot, err = h.rec.Vote().UpdateHash(vot, hsh)
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
	// Update the vote summary for the referenced posts.
	//

	var pos []*poststorage.Object
	{
		pos, err = h.rec.Post().UpdateVotes(vot)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != 0 {
		err = h.sto.Post().UpdatePost(pos)
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
