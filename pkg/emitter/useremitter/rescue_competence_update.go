package useremitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (r *Rescue) CompetenceUpdate(bal objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Core: &task.Core{
				task.Cancel: "10",
			},
			Meta: &task.Meta{
				objectlabel.ClaimObject: bal.String(),
				objectlabel.TaskOrigin:  objectlabel.OriginIntern,
				objectlabel.TaskWorker:  "competenceupdatehandler",
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
