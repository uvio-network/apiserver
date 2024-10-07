package claimemitter

import (
	"strconv"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (r *Rescue) ResolveCreate(blc uint64, pod objectid.ID) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Core: &task.Core{
				task.Cancel: "10",
			},
			Meta: &task.Meta{
				objectlabel.ClaimBlock:  strconv.FormatUint(blc, 10),
				objectlabel.ClaimObject: pod.String(),
				objectlabel.TaskOrigin:  objectlabel.OriginIntern,
				objectlabel.TaskWorker:  "resolvecreatehandler",
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
