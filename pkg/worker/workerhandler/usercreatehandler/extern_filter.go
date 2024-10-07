package usercreatehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *ExternHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.UserObject: "*",
		objectlabel.TaskOrigin: objectlabel.OriginExtern,
		objectlabel.TaskWorker: "usercreatehandler",
	})
}
