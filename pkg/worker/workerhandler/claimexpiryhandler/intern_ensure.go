package claimexpiryhandler

import (
	"context"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var dis []*poststorage.Object
	{
		dis, err = h.sto.Post().SearchExpiry(objectlabel.LifecycleDispute)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var pro []*poststorage.Object
	{
		pro, err = h.sto.Post().SearchExpiry(objectlabel.LifecyclePropose)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var res []*poststorage.Object
	{
		res, err = h.sto.Post().SearchExpiry(objectlabel.LifecycleResolve)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// It may very well happen that there are no claims to be expired. In that
	// case we stop processing here and prevent incurring further load on our RPC
	// provider.
	if len(dis) == 0 && len(pro) == 0 && len(res) == 0 {
		return nil
	}

	var blc uint64
	{
		blc, err = h.con.Client().BlockNumber(context.Background())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expireDispute(dis, blc, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expirePropose(pro, blc, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expireResolve(res, blc, bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) expireDispute(dis []*poststorage.Object, blc uint64, bud *budget.Budget) error {
	var err error

	for _, x := range dis[:bud.Claim(len(dis))] {
		{
			err = h.emi.Claim().Create(blc, x.ID, objectlabel.LifecycleResolve)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.sto.Post().DeleteExpiry([]*poststorage.Object{x})
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (h *InternHandler) expirePropose(pro []*poststorage.Object, blc uint64, bud *budget.Budget) error {
	var err error

	for _, x := range pro[:bud.Claim(len(pro))] {
		{
			err = h.emi.Claim().Create(blc, x.ID, objectlabel.LifecycleResolve)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.sto.Post().DeleteExpiry([]*poststorage.Object{x})
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (h *InternHandler) expireResolve(res []*poststorage.Object, blc uint64, bud *budget.Budget) error {
	var err error

	for _, x := range res[:bud.Claim(len(res))] {
		{
			err = h.emi.Claim().Update(blc, x.ID, objectlabel.LifecycleBalance)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.sto.Post().DeleteExpiry([]*poststorage.Object{x})
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
