package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Update(ctx context.Context, req *post.UpdateI) (*post.UpdateO, error) {
	// var err error

	// var use objectid.ID
	// {
	// 	use = userid.FromContext(ctx)
	// }

	// var out []objectstate.String

	//
	// Update by JSON-Patch
	//

	// {
	// 	var des []objectid.ID
	// 	var pat [][]*poststorage.Patch
	// 	for _, x := range req.Object {
	// 		if x.Intern != nil && x.Update != nil && x.Intern.Desc != "" {
	// 			des = append(des, objectid.ID(x.Intern.Desc))
	// 			pat = append(pat, inpPat(x.Update))
	// 		}
	// 	}

	// 	if len(des) != 0 {
	// 		var inp []*poststorage.Object
	// 		{
	// 			inp, err = h.des.SearchDesc(use, des)
	// 			if err != nil {
	// 				return nil, tracer.Mask(err)
	// 			}
	// 		}

	// 		//
	// 		// Verify the given input.
	// 		//

	// 		{
	// 			err = h.updateVrfyPtch(ctx, inp)
	// 			if err != nil {
	// 				return nil, tracer.Mask(err)
	// 			}
	// 		}

	// 		//
	// 		// Update the given resources.
	// 		//

	// 		{
	// 			lis, err := h.des.UpdatePtch(inp, pat)
	// 			if err != nil {
	// 				return nil, tracer.Mask(err)
	// 			}

	// 			out = append(out, lis...)
	// 		}
	// 	}
	// }

	//
	// Construct the RPC response.
	//

	var res *post.UpdateO
	{
		res = &post.UpdateO{}
	}

	// for _, x := range out {
	// 	res.Object = append(res.Object, &post.UpdateO_Object{
	// 		Intern: &post.UpdateO_Object_Intern{
	// 			Status: "",
	// 		},
	// 		Public: &post.UpdateO_Object_Public{},
	// 	})
	// }

	return res, nil
}
