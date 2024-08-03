package votehandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *vote.SearchI) (*vote.SearchO, error) {
	var out []*votestorage.Object

	var ids []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" {
			ids = append(ids, objectid.ID(x.Intern.Id))
		}
	}

	//
	// Search posts by ID.
	//

	if len(ids) != 0 {
		lis, err := h.sto.SearchVote(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct the RPC response.
	//

	var res *vote.SearchO
	{
		res = &vote.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "vote",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		res.Object = append(res.Object, &vote.SearchO_Object{
			Extern: []*vote.SearchO_Object_Extern{},
			Intern: &vote.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
				Owner:   x.Owner.String(),
			},
			Public: &vote.SearchO_Object_Public{
				Claim:  x.Claim.String(),
				Kind:   x.Kind,
				Option: converter.BoolToString(x.Option),
				Value:  converter.FloatToString(x.Value),
			},
		})
	}

	return res, nil
}
