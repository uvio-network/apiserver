package claimexpiryhandler

import (
	"context"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	oneWeek = time.Hour * 24 * 7
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
		var tre poststorage.Slicer
		{
			tre, err = h.sto.Post().SearchTree([]objectid.ID{x.Tree})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var res *poststorage.Object
		var fin bool
		{
			res, fin = finTre(time.Now().UTC(), tre)
		}

		// If the tree of the given resolve cannot be finalized, then we cannot go
		// ahead to emit a task for updating user balances. The resolve on which we
		// continue here may or may not be the last one that will ever be created.
		if !fin {
			continue
		}

		// We want to emit the task for updating user balances once, and only once.
		// We do that when we reconcile the very last resolve of any given claim
		// tree that can in fact be finalized.
		if x.ID == res.ID {
			err = h.emi.Claim().Update(blc, x.ID, objectlabel.LifecycleBalance)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// At this point the given claim tree can be finalized and we processed all
		// of the relevant resolves eventually. The very last step is then to remove
		// the processed resolves from the expiry queue.
		{
			err = h.sto.Post().DeleteExpiry([]*poststorage.Object{x})
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func finTre(now time.Time, tre poststorage.Slicer) (*poststorage.Object, bool) {
	var res poststorage.Slicer
	{
		res = tre.ObjectLifecycle(objectlabel.LifecycleResolve)
	}

	var lat *poststorage.Object
	{
		lat = res.LatestClaim()
	}

	// If there is not a single resolve, then this tree cannot be considered final.
	if lat == nil {
		return nil, false
	}

	// Any given claim tree is considered final if we have the maximum amount of 3
	// resolves at hand. This is because every originally proposed claim can be
	// disputed a maximum amount of two times. Then there is one resolve for the
	// original propose, and two resolves for every dispute after that.
	if len(res) == 3 && now.After(lat.Expiry) {
		return lat, true
	}

	// Any given claim tree is considered final if the latest resolve is outside
	// of its challenge window, which is its own expiry plus one additional week.
	// There can be only a single resolve for any given claim tree. In any event,
	// after this resolve moved out of its challenge window, no more disputes can
	// be created, and the latest resolution at hand becomes definitive and
	// binding.
	if now.After(lat.Expiry.Add(oneWeek)) {
		return lat, true
	}

	return nil, false
}
