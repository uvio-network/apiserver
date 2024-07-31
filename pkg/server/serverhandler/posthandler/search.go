package posthandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
	var out []*poststorage.Object

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var ids []objectid.ID
	var lab [][]string
	var tre []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil && x.Intern.Id != "" {
			ids = append(ids, objectid.ID(x.Intern.Id))
		}
		if x.Public != nil && x.Public.Labels != "" {
			var lis []string
			{
				lis = converter.StringToSlice(x.Public.Labels)
			}

			if len(lis) != 0 {
				lab = append(lab, lis)
			}
		}
		if x.Intern != nil && x.Intern.Tree != "" {
			tre = append(tre, objectid.ID(x.Intern.Tree))
		}
	}

	//
	// Search posts by ID.
	//

	if len(ids) != 0 {
		lis, err := h.pos.SearchPost(use, ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by labels.
	//

	if len(lab) != 0 {
		lis, err := h.pos.SearchLabels(use, lab)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by tree.
	//

	if len(tre) != 0 {
		lis, err := h.pos.SearchTree(use, tre)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Construct the RPC response.
	//

	var res *post.SearchO
	{
		res = &post.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "post",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		res.Object = append(res.Object, &post.SearchO_Object{
			Extern: []*post.SearchO_Object_Extern{},
			Intern: &post.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
				Owner:   x.Owner.String(),
				Tree:    x.Tree.String(),
			},
			Public: &post.SearchO_Object_Public{
				Expiry:    converter.TimeToString(x.Expiry),
				Kind:      x.Kind,
				Labels:    converter.SliceToString(x.Labels),
				Lifecycle: x.Lifecycle,
				Option:    converter.BoolToString(x.Option),
				Stake:     converter.FloatToString(x.Stake),
				Parent:    x.Parent.String(),
				Text:      x.Text,
				Token:     x.Token,
			},
		})
	}

	return res, nil
}
