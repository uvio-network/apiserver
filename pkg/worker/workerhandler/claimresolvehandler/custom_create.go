package claimresolvehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Create() *task.Task {
	return &task.Task{
		Cron: &task.Cron{
			task.Aevery: "hour",
		},
		Meta: &task.Meta{
			objectlabel.ClaimAction: objectlabel.ActionResolve,
			objectlabel.ClaimOrigin: objectlabel.OriginSystem,
		},
	}
}
