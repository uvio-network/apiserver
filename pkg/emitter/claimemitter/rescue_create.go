package claimemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (r *Rescue) Create(cla objectid.ID, lif objectlabel.DesiredLifecycle) error {
	var tas *task.Task
	{
		tas = &task.Task{
			Meta: &task.Meta{
				objectlabel.ClaimAction:    objectlabel.ActionCreate,
				objectlabel.ClaimLifecycle: string(lif),
				objectlabel.ClaimObject:    cla.String(),
				objectlabel.ClaimOrigin:    objectlabel.OriginIntern,
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