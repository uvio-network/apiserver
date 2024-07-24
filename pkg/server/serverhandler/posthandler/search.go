package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
	// var out []*poststorage.Object

	// var use objectid.ID
	// {
	// 	use = userid.FromContext(ctx)
	// }

	// var evn []objectid.ID
	// for _, x := range req.Object {
	// 	if x.Public != nil && x.Public.Evnt != "" {
	// 		evn = append(evn, objectid.ID(x.Public.Evnt))
	// 	}
	// }

	//
	// Search the given resources.
	//

	// {
	// 	lis, err := h.pos.SearchPost(use, pos)
	// 	if err != nil {
	// 		return nil, tracer.Mask(err)
	// 	}

	// 	out = append(out, lis...)
	// }

	//
	// Construct the RPC response.
	//

	var res *post.SearchO
	{
		res = &post.SearchO{}
	}

	// if limiter.Log(len(out)) {
	// 	h.log.Log(
	// 		context.Background(),
	// 		"level", "warning",
	// 		"message", "search response truncated",
	// 		"limit", strconv.Itoa(limiter.Default),
	// 		"resource", "post",
	// 		"total", strconv.Itoa(len(out)),
	// 	)
	// }

	// for _, x := range out[:limiter.Len(len(out))] {
	// posts marked to be deleted cannot be searched anymore.
	// if !x.Dltd.IsZero() {
	// 	continue
	// }

	// 	res.Object = append(res.Object, &post.SearchO_Object{
	// 		Extern: []*post.SearchO_Object_Extern{},
	// 		Intern: &post.SearchO_Object_Intern{
	// 			Created: "",
	// 			Post:    "",
	// 			User:    "",
	// 		},
	// 		Public: &post.SearchO_Object_Public{
	// 			Text: "",
	// 		},
	// 	})
	// }

	return res, nil
}
