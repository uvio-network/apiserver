package uvxminthandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// var err error

	var use objectid.ID
	{
		use = objectid.ID(tas.Meta.Get(objectlabel.UserObject))
	}

	fmt.Printf("%#v\n", use)

	// TODO mint UVX onchain

	return nil
}
