package posthandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
	var err error
	var out []*poststorage.Object

	var use objectid.ID
	{
		use = userid.FromContext(ctx)
	}

	var ids []objectid.ID
	var lab [][]string
	var pag []int
	var tim bool
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

		if x.Symbol != nil && x.Symbol.Time == "latest" {
			{
				tim = true
			}

			if req.Filter != nil && req.Filter.Paging != nil && req.Filter.Paging.Kind == "page" {
				pag, err = createPage(req.Filter.Paging)
				if err != nil {
					return nil, tracer.Mask(err)
				}
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
		lis, err := h.sto.SearchPost(use, ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by labels.
	//

	if len(lab) != 0 {
		lis, err := h.sto.SearchLabels(use, lab)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by time.
	//

	if tim && len(pag) == 2 {
		lis, err := h.sto.SearchPage(use, pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by tree.
	//

	if len(tre) != 0 {
		lis, err := h.sto.SearchTree(use, tre)
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

func createPage(pag *post.SearchI_Filter_Paging) ([]int, error) {
	if pag.Start == "" {
		return nil, tracer.Maskf(runtime.QueryPagingMissingError, "Paging.Start")
	}

	if pag.Stop == "" {
		return nil, tracer.Maskf(runtime.QueryPagingMissingError, "Paging.Stop")
	}

	beg, err := strconv.Atoi(pag.Start)
	if err != nil {
		return nil, tracer.Maskf(runtime.QueryPagingInvalidError, "Paging.Start")
	}

	end, err := strconv.Atoi(pag.Stop)
	if err != nil {
		return nil, tracer.Maskf(runtime.QueryPagingInvalidError, "Paging.Stop")
	}

	if beg < 0 {
		return nil, tracer.Maskf(runtime.QueryPagingNegativeError, "Paging.Start")
	}

	if end < 0 {
		return nil, tracer.Maskf(runtime.QueryPagingNegativeError, "Paging.Stop")
	}

	ran := end - beg
	if ran < 1 || ran > 1000 {
		return nil, tracer.Mask(runtime.QueryPagingRangeError)
	}

	return []int{beg, end}, nil
}
