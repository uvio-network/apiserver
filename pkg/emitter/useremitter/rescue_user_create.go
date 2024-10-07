package useremitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (r *Rescue) UserCreate(use objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.UserObject: use.String(),
				objectlabel.TaskOrigin: objectlabel.OriginExtern,
				objectlabel.TaskWorker: "usercreatehandler",
			},
		}
	}

	{
		err := r.res.Create(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
