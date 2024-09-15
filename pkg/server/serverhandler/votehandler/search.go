package votehandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *vote.SearchI) (*vote.SearchO, error) {
	var out []*votestorage.Object

	var cla []objectid.ID
	var own []objectid.ID
	var vot []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" {
			vot = append(vot, objectid.ID(x.Intern.Id))
		}
		if x.Intern != nil && x.Intern.Owner != "" {
			if x.Intern.Owner == "self" {
				own = append(own, userid.FromContext(ctx))
			} else {
				own = append(own, objectid.ID(x.Intern.Owner))
			}
		}
		if x.Public != nil && x.Public.Claim != "" {
			cla = append(cla, objectid.ID(x.Public.Claim))
		}
	}

	//
	// Search votes by claim ID.
	//

	if len(cla) != 0 && len(own) == 0 {
		lis, err := h.sto.Vote().SearchClaim(cla)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search votes by user ID.
	//

	if len(own) != 0 && len(cla) == 0 {
		lis, err := h.sto.Vote().SearchOwner(own)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search votes by claim ID and user ID.
	//

	if len(cla) != 0 && len(own) != 0 {
		lis, err := h.sto.Vote().SearchOwnerClaim(own, cla)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search votes by vote ID.
	//

	if len(vot) != 0 {
		lis, err := h.sto.Vote().SearchVote(vot)
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
			Intern: &vote.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
				Owner:   x.Owner.String(),
			},
			Public: &vote.SearchO_Object_Public{
				Chain:     x.Chain,
				Claim:     x.Claim.String(),
				Hash:      converter.SliceToString(x.Lifecycle.Hash),
				Kind:      x.Kind,
				Lifecycle: x.Lifecycle.String(),
				Meta:      x.Meta,
				Option:    converter.BoolToString(x.Option),
				Value:     converter.FloatToString(x.Value),
			},
		})
	}

	return res, nil
}
