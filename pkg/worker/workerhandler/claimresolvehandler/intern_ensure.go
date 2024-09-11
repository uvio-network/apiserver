package claimresolvehandler

import (
	"fmt"
	"math/big"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var cla []*poststorage.Object
	{
		cla, err = h.sto.Post().SearchExpiry(objectlabel.LifecyclePropose)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range cla {
		// TODO create resolve claims onchain
		fmt.Printf("cla %#v\n", big.NewInt(x.ID.Int()))
	}

	return nil
}
