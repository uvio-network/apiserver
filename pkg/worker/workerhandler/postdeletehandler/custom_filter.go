package postdeletehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *CustomHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.PostAction: objectlabel.ActionDelete,
		objectlabel.PostObject: "*",
		objectlabel.PostOrigin: objectlabel.OriginCustom,
	})
}
