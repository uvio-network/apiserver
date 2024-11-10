package notehandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/note"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *note.SearchI) (*note.SearchO, error) {
	var err error
	var out notestorage.Slicer

	var kin []string
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Kind != "" {
			kin = append(kin, x.Public.Kind)
		}
	}

	var pag []int64
	var tim []int64
	if req.Filter != nil && req.Filter.Paging != nil {
		if req.Filter.Paging.Kind == "page" {
			pag, err = createPage(req.Filter.Paging)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
		if req.Filter.Paging.Kind == "time" {
			tim, err = createPage(req.Filter.Paging)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	//
	// Search for the calling user object.
	//

	var use []*userstorage.Object
	{
		use, err = h.sto.User().SearchUser([]objectid.ID{userid.FromContext(ctx)})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Search notes by page.
	//

	if len(pag) != 0 {
		lis, err := h.sto.Note().SearchPage(use[0].ID, int(pag[0]), int(pag[1]))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search notes by time.
	//

	if len(tim) != 0 {
		lis, err := h.sto.Note().SearchTime(use[0].ID, tim[0], tim[1])
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Filter the found note objects.
	//

	{
		out = out.ObjectKind(kin)
	}

	//
	// Construct the RPC response.
	//

	var res *note.SearchO
	{
		res = &note.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "note",
			"total", strconv.Itoa(len(out)),
		)
	}

	if len(use) == 1 {
		res.Filter = &note.SearchO_Filter{
			Pointer: use[0].Pointer,
		}
	}

	for _, x := range out[:limiter.Len(len(out))] {
		res.Object = append(res.Object, &note.SearchO_Object{
			Intern: &note.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
				Owner:   x.Owner.String(),
			},
			Public: &note.SearchO_Object_Public{
				Kind:     x.Kind,
				Message:  x.Message,
				Resource: x.Resource.String(),
			},
		})
	}

	return res, nil
}

func createPage(pag *note.SearchI_Filter_Paging) ([]int64, error) {
	beg, err := converter.PageToInt(pag.Start)
	if err != nil {
		return nil, tracer.Maskf(runtime.QueryPagingInvalidError, "Paging.Start")
	}

	end, err := converter.PageToInt(pag.Stop)
	if err != nil {
		return nil, tracer.Maskf(runtime.QueryPagingInvalidError, "Paging.Stop")
	}

	ran := end - beg
	if ran < 1 {
		return nil, tracer.Mask(runtime.QueryPagingTimeError)
	}

	return []int64{beg, end}, nil
}
