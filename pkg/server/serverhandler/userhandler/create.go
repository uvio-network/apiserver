package userhandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/server/context/subjectclaim"
	"github.com/uvio-network/apiserver/pkg/server/context/usercreated"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
	var err error

	var img string
	var nam string
	if req.Object[0].Public != nil {
		img = req.Object[0].Public.Image
		nam = req.Object[0].Public.Name
	}

	var inp *userstorage.Object
	{
		inp = &userstorage.Object{
			Image: objectfield.String{
				Data: img,
			},
			Name: objectfield.String{
				Data: nam,
			},
			Subject: []string{subjectclaim.FromContext(ctx)},
		}
	}

	//
	// Create the given resource.
	//

	var out *userstorage.Object
	{
		out, err = h.sto.User().Create(inp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Emit an event for new users being created.
	//

	if usercreated.FromContext(ctx) {
		err = h.emi.User().UserCreate(out.ID)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *user.CreateO
	{
		res = &user.CreateO{
			Object: []*user.CreateO_Object{
				{
					Intern: &user.CreateO_Object_Intern{
						Created: converter.TimeToString(out.Created),
						Id:      out.ID.String(),
					},
				},
			},
		}
	}

	return res, nil
}
