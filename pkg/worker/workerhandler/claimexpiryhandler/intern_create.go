package claimexpiryhandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			task.Aevery: "hour",
		},
		Meta: &task.Meta{
			objectlabel.ClaimObject: "*",
			objectlabel.TaskOrigin:  objectlabel.OriginIntern,
			objectlabel.TaskWorker:  "claimexpiryhandler",
		},
	}
}
