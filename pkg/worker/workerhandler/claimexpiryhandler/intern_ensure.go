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

	{
		err = h.expirePropose(bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.expireResolve(bud)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) expirePropose(bud *budget.Budget) error {
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
			err = h.emi.Claim().Create(x.ID, objectlabel.LifecycleResolve)
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

func (h *InternHandler) expireResolve(bud *budget.Budget) error {
	var err error

	var cla []*poststorage.Object
	{
		cla, err = h.sto.Post().SearchExpiry(objectlabel.LifecycleResolve)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range cla[:bud.Claim(len(cla))] {
		var tre []*poststorage.Object
		{
			tre, err = h.sto.Post().SearchTree([]objectid.ID{x.ID})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var res []*poststorage.Object
		{
			res = searchResolve(tre)
		}

		// It happens that a resolve has expired, while other claims supersede it
		// within a tree. That means if the latest claim in a tree is not a resolve,
		// then we cannot just yet update balances, because one of two possible
		// scenarios can be present. Either somebody made use of the dispute
		// challenge, or the superseding claim is of lifecycle phase balance. The
		// former means that the resolution is still ongoing. The latter means that
		// balances have already been updated.
		if len(res) == 0 {
			h.log.Log(
				context.Background(),
				"level", "debug",
				"message", "resolve expired but balances cannot be updated just yet",
				"resolve", x.ID.String(),
			)

			continue
		}

		{
			res = filterExpiry(time.Now(), res)
		}

		// We are given all resolves within a tree. Once a resolve expired, there is
		// a 7 day dispute challenge on the latest resolve, within any given tree,
		// that we have to wait for. Therefore we can either forward all resolves,
		// or none, which means that we have to filter out all claims that are still
		// inside of their respective dispute challenge.
		if len(res) == 0 {
			continue
		}

		for _, y := range res {
			{
				err = h.emi.Claim().Update(y.Parent, objectlabel.LifecycleBalance)
				if err != nil {
					return tracer.Mask(err)
				}
			}

			{
				err = h.sto.Post().DeleteExpiry([]*poststorage.Object{y})
				if err != nil {
					return tracer.Mask(err)
				}
			}
		}
	}

	return nil
}

func filterExpiry(now time.Time, res poststorage.Slicer) []*poststorage.Object {
	var lat *poststorage.Object
	{
		lat = res.LatestClaim()
	}

	if lat != nil && lat.Expiry.After(now.Add(oneWeek)) {
		return res
	}

	return nil
}

func searchResolve(tre poststorage.Slicer) []*poststorage.Object {
	var lat *poststorage.Object
	{
		lat = tre.LatestClaim()
	}

	if lat != nil && lat.Lifecycle.Is(objectlabel.LifecycleResolve) {
		return tre.ObjectLifecycle(objectlabel.LifecycleResolve)
	}

	return nil
}
