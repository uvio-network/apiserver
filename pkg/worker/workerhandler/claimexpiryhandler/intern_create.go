package claimexpiryhandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			task.Aevery: "minute",
		},
		Meta: &task.Meta{
			objectlabel.ClaimAction: objectlabel.ActionExpiry,
			objectlabel.ClaimObject: "*",
			objectlabel.ClaimOrigin: objectlabel.OriginIntern,
		},
	}
}
