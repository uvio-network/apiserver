package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Create(ctx context.Context, req *post.CreateI) (*post.CreateO, error) {
	// var err error

	// var inp []*poststorage.Object
	// for _, x := range req.Object {
	// 	inp = append(inp, &poststorage.Object{
	// 		User: userid.FromContext(ctx),
	// 	})
	// }

	//
	// Verify the given input.
	//

	//
	// Create the given resources.
	//

	// var out []*poststorage.Object

	//
	// Construct the RPC response.
	//

	var res *post.CreateO
	{
		res = &post.CreateO{}
	}

	// for _, x := range out {
	// 	res.Object = append(res.Object, &post.CreateO_Object{
	// 		Intern: &post.CreateO_Object_Intern{
	// 			Created: "",
	// 			Post:    "",
	// 		},
	// 	})
	// }

	return res, nil
}
