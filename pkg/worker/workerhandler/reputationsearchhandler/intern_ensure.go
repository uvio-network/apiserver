package reputationsearchhandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	// TODO search for privileged voter addresses

	var blc uint64
	var dis objectid.ID

	{
		err = h.emi.Claim().ResolveCreate(blc, dis)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
