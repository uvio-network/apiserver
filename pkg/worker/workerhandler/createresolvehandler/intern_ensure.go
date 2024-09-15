package createresolvehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// var err error

	var cla objectid.ID
	{
		cla = objectid.ID(tas.Meta.Get(objectlabel.ClaimObject))
	}

	fmt.Printf("cla %#v\n", cla)

	// TODO create post object with lifecycle phase resolve
	// TODO create resolve onchain
	// TODO update hash in post object with transaction hash once confirmed onchain

	return nil
}
