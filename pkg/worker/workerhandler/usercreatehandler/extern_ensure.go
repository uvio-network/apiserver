package usercreatehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *ExternHandler) Ensure(tas *task.Task) error {
	var err error

	var use objectid.ID
	{
		use = objectid.ID(tas.Meta.Get(objectlabel.UserObject))
	}

	{
		err = h.emi.UVX().UvxMint(use)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
