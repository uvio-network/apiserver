package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Delete(ctx context.Context, req *post.DeleteI) (*post.DeleteO, error) {
	// var err error

	// var use objectid.ID
	// {
	// 	use = userid.FromContext(ctx)
	// }

	// var inp []*poststorage.Object
	// {
	// 	inp, err = h.pos.SearchPost(use, pos)
	// 	if err != nil {
	// 		return nil, tracer.Mask(err)
	// 	}
	// }

	//
	// Verify the given input.
	//

	//
	// Delete the given resources.
	//

	// var out []objectstate.String

	//
	// Construct the RPC response.
	//

	var res *post.DeleteO
	{
		res = &post.DeleteO{}
	}

	// for _, x := range out {
	// 	res.Object = append(res.Object, &post.DeleteO_Object{
	// 		Intern: &post.DeleteO_Object_Intern{
	// 			Status: "",
	// 		},
	// 		Public: &post.DeleteO_Object_Public{},
	// 	})
	// }

	return res, nil
}
