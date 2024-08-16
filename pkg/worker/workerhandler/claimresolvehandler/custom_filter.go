package claimresolvehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.ClaimAction: objectlabel.ActionResolve,
		objectlabel.ClaimOrigin: objectlabel.OriginSystem,
	})
}
