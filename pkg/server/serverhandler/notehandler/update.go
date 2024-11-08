package notehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/note"
	"github.com/uvio-network/apiserver/pkg/object/objectstatus"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) Update(ctx context.Context, req *note.UpdateI) (*note.UpdateO, error) {
	var err error

	var kin []string
	var poi []string
	for _, x := range req.Object {
		if x.Public != nil && x.Public.Kind != "" && x.Public.Pointer != "" {
			kin = append(kin, x.Public.Kind)
			poi = append(poi, x.Public.Pointer)
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
	// Update the note pointers.
	//

	for i := range use {
		for j := range kin {
			use[i].Pointer[kin[j]] = poi[j]
		}
	}

	//
	// Update the given user.
	//

	{
		err = h.sto.User().UpdateUser(use)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	//
	// Construct the RPC response.
	//

	var res *note.UpdateO
	{
		res = &note.UpdateO{}
	}

	for range use {
		res.Object = append(res.Object, &note.UpdateO_Object{
			Intern: &note.UpdateO_Object_Intern{
				Status: objectstatus.Updated,
			},
		})
	}

	return res, nil
}