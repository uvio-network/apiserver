package noteemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (r *Rescue) NoteCreate(kin string, mes string, res objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.NoteKind:     kin,
				objectlabel.NoteMessage:  mes,
				objectlabel.TaskOrigin:   objectlabel.OriginIntern,
				objectlabel.NoteResource: res.String(),
				objectlabel.TaskWorker:   "notecreatehandler",
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
