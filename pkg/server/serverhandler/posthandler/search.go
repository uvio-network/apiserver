package posthandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
	var err error
	var out poststorage.Slicer

	var ids []objectid.ID
	var lab [][]string
	var lif []objectlabel.DesiredLifecycle
	var own []objectid.ID
	var tim bool
	var tre []objectid.ID
	var vot []objectid.ID
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

		if x.Public != nil && x.Public.Lifecycle != "" {
			lif = append(lif, objectlabel.DesiredLifecycle(x.Public.Lifecycle))
		}

		if x.Intern != nil && x.Intern.Owner != "" {
			own = append(own, objectid.ID(x.Intern.Owner))
		}

		if x.Symbol != nil && x.Symbol.Time == "page" {
			tim = true
		}

		if x.Symbol != nil && x.Symbol.Vote != "" {
			if x.Symbol.Vote == "self" {
				vot = append(vot, userid.FromContext(ctx))
			} else {
				vot = append(vot, objectid.ID(x.Symbol.Vote))
			}
		}

		if x.Intern != nil && x.Intern.Tree != "" {
			tre = append(tre, objectid.ID(x.Intern.Tree))
		}
	}

	var pag []int
	if req.Filter != nil && req.Filter.Paging != nil && req.Filter.Paging.Kind == "page" {
		pag, err = createPage(req.Filter.Paging)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Search posts by ID. If the posts being searched turn out to be of kind
	// "claim", then do also search for all of their comments, if any.
	//

	if len(ids) != 0 {
		lis, err := h.sto.Post().SearchPost(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by labels.
	//

	if len(lab) != 0 {
		lis, err := h.sto.Post().SearchLabel(lab)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by lifecycle.
	//

	if len(lif) != 0 {
		lis, err := h.sto.Post().SearchLifecycle(lif)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by owner.
	//

	if len(own) != 0 {
		lis, err := h.sto.Post().SearchOwner(own)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by creation time.
	//

	if tim && len(pag) == 2 {
		lis, err := h.sto.Post().SearchPage(objectlabel.LifecycleCreated, pag[0], pag[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by tree.
	//

	if len(tre) != 0 {
		lis, err := h.sto.Post().SearchTree(tre)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search posts by the voting users.
	//

	if len(vot) != 0 {
		var sli votestorage.Slicer
		{
			sli, err = h.sto.Vote().SearchOwner(vot)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		var cla []objectid.ID
		{
			cla = generic.Unique(sli.Claim())
		}

		if len(cla) != 0 {
			lis, err := h.sto.Post().SearchPost(cla)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	var all []objectid.ID
	{
		all = out.Tree()
	}

	// As a final step return the entire trees of anything that we found above.
	if len(all) != 0 {
		out, err = h.sto.Post().SearchTree(all)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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
			Extern: &post.SearchO_Object_Extern{
				Samples: x.Samples,
			},
			Intern: &post.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
				Owner:   x.Owner.String(),
				Tree:    x.Tree.String(),
			},
			Public: &post.SearchO_Object_Public{
				Chain:     x.Chain,
				Contract:  x.Contract,
				Expiry:    converter.TimeToString(x.Expiry),
				Hash:      converter.SliceToString(x.Lifecycle.Hash),
				Kind:      x.Kind,
				Labels:    converter.SliceToString(x.Labels),
				Lifecycle: x.Lifecycle.String(),
				Meta:      x.Meta,
				Parent:    x.Parent.String(),
				Text:      x.Text,
				Token:     x.Token,
				Summary:   converter.FloatsToString(x.Summary),
			},
		})
	}

	return res, nil
}

func createPage(pag *post.SearchI_Filter_Paging) ([]int, error) {
	beg, err := converter.PageToInt(pag.Start)
	if err != nil {
		return nil, tracer.Maskf(runtime.QueryPagingInvalidError, "Paging.Start")
	}

	end, err := converter.PageToInt(pag.Stop)
	if err != nil {
		return nil, tracer.Maskf(runtime.QueryPagingInvalidError, "Paging.Stop")
	}

	ran := end - beg
	if ran < 1 || ran > 1000 {
		return nil, tracer.Mask(runtime.QueryPagingPageError)
	}

	return []int{int(beg), int(end)}, nil
}
