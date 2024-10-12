package userhandler

import (
	"context"
	"strconv"

	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/server/context/subjectclaim"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/server/limiter"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	var out []*userstorage.Object

	var slf bool
	var use []objectid.ID
	for _, x := range req.Object {
		if x.Intern != nil {
			if x.Intern.Id == "self" {
				slf = true
			} else if x.Intern.Id != "" {
				use = append(use, objectid.ID(x.Intern.Id))
			}
		}
	}

	//
	// Search users by ID.
	//

	if len(use) != 0 {
		lis, err := h.sto.User().SearchUser(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	//
	// Search users by subject claim.
	//

	if slf {
		var sub string
		{
			sub = subjectclaim.FromContext(ctx)
		}

		{
			obj, err := h.sto.User().SearchSubject(sub)
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, obj)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *user.SearchO
	{
		res = &user.SearchO{}
	}

	if limiter.Log(len(out)) {
		h.log.Log(
			context.Background(),
			"level", "warning",
			"message", "search response truncated",
			"limit", strconv.Itoa(limiter.Default),
			"resource", "user",
			"total", strconv.Itoa(len(out)),
		)
	}

	for _, x := range out[:limiter.Len(len(out))] {
		res.Object = append(res.Object, &user.SearchO_Object{
			Intern: &user.SearchO_Object_Intern{
				Created: converter.TimeToString(x.Created),
				Id:      x.ID.String(),
			},
			Public: &user.SearchO_Object_Public{
				Image: x.Image.Data,
				Name:  x.Name.Data,
			},
		})
	}

	return res, nil
}
