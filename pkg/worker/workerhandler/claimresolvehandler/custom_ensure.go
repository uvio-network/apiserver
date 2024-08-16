package claimresolvehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	fmt.Printf("henlo\n")
	// TODO
	//
	//     find claims with lifecycle "propose" that are expired
	//
	//     create new claims with lifecycle "resolve" and same tree ID
	//

	return nil
}
