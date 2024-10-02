package balanceupdatehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Filter(tas *task.Task) bool {
	// The tasks filtered for here are emitted using Emitter.Claim.Update in the
	// file below.
	//
	//     pkg/worker/workerhandler/claimexpiryhandler/intern_ensure.go
	//
	return tas.Meta.Has(map[string]string{
		objectlabel.ClaimAction:    objectlabel.ActionUpdate,
		objectlabel.ClaimBlock:     "*",
		objectlabel.ClaimLifecycle: string(objectlabel.LifecycleBalance),
		objectlabel.ClaimObject:    "*",
		objectlabel.ClaimOrigin:    objectlabel.OriginIntern,
	})
}
