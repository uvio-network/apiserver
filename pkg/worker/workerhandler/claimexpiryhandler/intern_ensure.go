package claimexpiryhandler

import (
	"context"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/objectid"
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
		var sec bool
		{
			sec, err = h.secDis(x)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// If x is the second dispute within this claim tree, then create a
		// ReputationSearch task instead of a ResolveCreate task. In the end we are
		// going to create the resolve anyway, but at the second dispute we are
		// injecting the task to search for privileged voters first, which in turn
		// will create the ResolveCreate task with additional meta data, if any.
		if sec {
			err = h.emi.User().ReputationSearch(x.ID)
			if err != nil {
				return tracer.Mask(err)
			}
		} else {
			err = h.emi.Claim().ResolveCreate(blc, x.ID)
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
			err = h.emi.Claim().ResolveCreate(blc, x.ID)
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
			err = h.emi.Claim().BalanceUpdate(blc, x.ID)
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

func (h *InternHandler) secDis(dis *poststorage.Object) (bool, error) {
	var err error

	var tre poststorage.Slicer
	{
		tre, err = h.sto.Post().SearchTree([]objectid.ID{dis.Tree})
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	var all poststorage.Slicer
	{
		all = tre.ObjectLifecycle(objectlabel.LifecycleDispute)
	}

	if len(all) < 1 {
		return false, tracer.Maskf(runtime.ExecutionFailedError, "expected at least one dispute for claim tree %s", dis.Tree)
	}

	var lat *poststorage.Object
	{
		lat = all.LatestObject()
	}

	// If there are 2 disputes in this claim tree, and if the given dispute is the
	// latest of those two disputes, then return true.
	return len(all) >= 2 && dis.ID == lat.ID, nil
}
