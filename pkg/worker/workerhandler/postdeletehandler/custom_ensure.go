package postdeletehandler

import (
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *CustomHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// var pos objectid.ID
	// {
	// 	pos = objectid.ID(tas.Meta.Get(objectlabel.PostObject))
	// }

	// {
	// 	err := h.deletePost(pos, bud)
	// 	if err != nil {
	// 		return tracer.Mask(err)
	// 	}
	// }

	return nil
}

// func (h *CustomHandler) deletePost(inp objectid.ID, bud *budget.Budget) error {
// var err error

// var pos []*poststorage.Object
// {
// 	pos, err = h.pos.SearchPost("", []objectid.ID{inp})
// 	if err != nil {
// 		return tracer.Mask(err)
// 	}
// }

// {
// 	_, err := h.pos.DeletePost(pos[:bud.Claim(len(pos))])
// 	if err != nil {
// 		return tracer.Mask(err)
// 	}
// }

// return nil
// }
