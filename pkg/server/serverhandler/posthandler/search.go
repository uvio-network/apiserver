package posthandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
	var err error
	var out []*poststorage.Object

	var ids []objectid.ID
	var lab [][]string
	var own []objectid.ID
	var pag []int
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

		if x.Intern != nil && x.Intern.Owner != "" {
			own = append(own, objectid.ID(x.Intern.Owner))
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

	//
	// Search posts by ID. If the posts being searched turn out to be of kind
	// "claim", then do also search for all of their comments, if any.
	//

	if len(ids) != 0 {
		var pos poststorage.Slicer
		{
			pos, err = h.sto.Post().SearchPost(ids)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, pos...)
		}

		var cla []objectid.ID
		{
			cla = poststorage.Slicer(pos).ObjectKind("claim").ID()
		}

		if len(cla) != 0 {
			com, err := h.sto.Post().SearchComment(cla)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, com...)
		}
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
		lis, err := h.sto.Post().SearchCreated(pag[0], pag[1])
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

	//
	// Search for all the parent claims referenced by comments that we have not
	// yet fetched. This last step is done to ensure any comment does also return
	// its referenced parent claim.
	//

	var par []objectid.ID
	var pos []objectid.ID
	{
		par = poststorage.Slicer(out).ObjectKind("comment").Parent()
		pos = poststorage.Slicer(out).ID()
	}

	var com []objectid.ID
	{
		com = generic.Select(pos, par)
	}

	if len(com) != 0 {
		lis, err := h.sto.Post().SearchPost(com)
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
				Chain:     x.Chain,
				Expiry:    converter.TimeToString(x.Expiry),
				Hash:      x.Hash,
				Kind:      x.Kind,
				Labels:    converter.SliceToString(x.Labels),
				Lifecycle: x.Lifecycle,
				Meta:      x.Meta,
				Parent:    x.Parent.String(),
				Text:      x.Text,
				Token:     x.Token,
				Votes:     converter.FloatsToString(x.Votes),
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
