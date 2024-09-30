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

	var blc uint64
	{
		blc, err = h.con.Client().BlockNumber(context.Background())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expirePropose(bud, blc)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expireResolve(bud, blc)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expireDispute(bud, blc)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) expireDispute(bud *budget.Budget, blc uint64) error {
	var err error

	var cla []*poststorage.Object
	{
		cla, err = h.sto.Post().SearchExpiry(objectlabel.LifecycleDispute)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range cla[:bud.Claim(len(cla))] {
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

func (h *InternHandler) expirePropose(bud *budget.Budget, blc uint64) error {
	var err error

	var cla []*poststorage.Object
	{
		cla, err = h.sto.Post().SearchExpiry(objectlabel.LifecyclePropose)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range cla[:bud.Claim(len(cla))] {
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

func (h *InternHandler) expireResolve(bud *budget.Budget, blc uint64) error {
	var err error

	var cla []*poststorage.Object
	{
		cla, err = h.sto.Post().SearchExpiry(objectlabel.LifecycleResolve)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range cla[:bud.Claim(len(cla))] {
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
