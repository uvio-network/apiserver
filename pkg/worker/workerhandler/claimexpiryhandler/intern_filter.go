package claimexpiryhandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Filter(tas *task.Task) bool {
	return tas.Meta.Has(map[string]string{
		objectlabel.ClaimAction: objectlabel.ActionExpiry,
		objectlabel.ClaimObject: "*",
		objectlabel.ClaimOrigin: objectlabel.OriginIntern,
	})
}
