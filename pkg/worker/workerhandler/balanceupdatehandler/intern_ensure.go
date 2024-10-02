package balanceupdatehandler

import (
	"context"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	// maxUser is the maximum amount of user balances to update at a time. This
	// number is a consideration of available block gas. If a claim were to have
	// 5000 market participants, then updating all user balances in a single
	// transaction may simply not possible. This parameter allows to chunk the
	// balance updating process according to the abilities of the underlying
	// system.
	maxUser uint64 = 25

	// maxWait is the maximum amount of time to wait for transactions per task
	// execution. We need to limit the task execution time because every worker
	// has only 30 seconds to process a task by default before ownership of the
	// provided task will be revoked. Therefore defining maxWait under 30 seconds
	// gives us enough time to gracefully process the task given to us.
	maxWait = 20 * time.Second
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var pod *poststorage.Object
	var res *poststorage.Object
	var bal *poststorage.Object
	{
		pod, res, bal, err = h.searchClaims(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cla claimscontract.Interface
	{
		cla = h.con.Claims(pod.Contract)
	}

	if tas.Sync == nil {
		tas.Sync = &task.Sync{}
	}

	// We are receiving the block number at the time of task creation and use it
	// as a paging pointer to retry task execution a couple of times.
	{
		tas.Sync.Set(task.Paging, tas.Meta.Get(objectlabel.ClaimBlock))
	}

	// If the next claim within this tree relative to the provided resolve is nil,
	// then that means that we have not yet created the required balance. And so
	// we can proceed to create the post object for the requested balance.
	if bal == nil {
		bal, err = h.rec.Post().CreateBalance(res)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Check whether the given claim has been settled onchain already.
	var stl bool
	{
		stl, err = cla.SearchResolve(pod.ID, claimscontract.CLAIM_BALANCE_S)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if !stl {
		var txn *types.Transaction
		{
			txn, err = cla.UpdateBalance(pod.ID, maxUser)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.con.Wait(txn, maxWait)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// After we have updated balances, check the claim state again, to see if
		// our most recent call to updateBalance has finalized the given claim
		// already.
		{
			stl, err = cla.SearchResolve(pod.ID, claimscontract.CLAIM_BALANCE_S)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// If the given claim has still not been settled, we stop processing here
		// and allow the task to be given back to us again later, since the task
		// still contains the block number at which we started processing in the
		// first place. That way we can chunk our work and iterate multiple times
		// until the claim could be fully settled onchain.
		if !stl {
			return nil
		}
	}

	if bal.Lifecycle.Pending() {
		var blc uint64
		{
			blc, err = ensInt(tas.Meta.Get(objectlabel.ClaimBlock))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var hsh []common.Hash
		{
			hsh, err = cla.BalanceUpdated(blc, pod.ID)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.rec.Post().UpdateBalance(bal, hsh)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	// Once the new post object got updated with all the associated transaction
	// hashes, we can remove the paging pointer and all sync state from the task
	// so that it can be deleted eventually by the rescue engine.
	{
		tas.Sync = nil
	}

	var blc uint64
	{
		blc, err = h.con.Client().BlockNumber(context.Background())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Finally emit the Settled events so that we can create notifications and
	// update user metrics.

	{
		err = h.emi.Claim().Create(blc, bal.ID, objectlabel.LifecycleSettled)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.emi.Claim().Update(blc, bal.ID, objectlabel.LifecycleSettled)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) searchClaims(tas *task.Task) (*poststorage.Object, *poststorage.Object, *poststorage.Object, error) {
	var err error

	// The claim ID obtained here is the ID of the resolve that expired when the
	// task that we are processing right now got emitted.
	var cla objectid.ID
	{
		cla = objectid.ID(tas.Meta.Get(objectlabel.ClaimObject))
	}

	var pos poststorage.Slicer
	{
		pos, err = h.sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, nil, nil, tracer.Mask(err)
		}
	}

	if len(pos) != 1 {
		return nil, nil, nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one post object for ID %s", cla)
	}

	// This is the claim with lifecycle phase resolve that has been expired.
	var res *poststorage.Object
	{
		res = pos[0]
	}

	var tre poststorage.Slicer
	{
		tre, err = h.sto.Post().SearchTree([]objectid.ID{res.Tree})
		if err != nil {
			return nil, nil, nil, tracer.Mask(err)
		}
	}

	// This is the claim here is either a propose or a dispute in first or second
	// instance. In either case, those are the claims that we are tasked to
	// settle.
	var pod *poststorage.Object
	{
		pod = tre.IDClaim(res.Parent)
	}

	var bal *poststorage.Object
	{
		bal, err = balTre(res.ID, tre)
		if err != nil {
			return nil, nil, nil, tracer.Mask(err)
		}
	}

	return pod, res, bal, nil
}

func balTre(res objectid.ID, tre poststorage.Slicer) (*poststorage.Object, error) {
	// We want to find the balance given the provided resolve "res". The "res"
	// here is effectively the parent of the balance that we want to find. The
	// "res" can also be the parent of comments, disputes and settles. Below we
	// search for the post objects that define the "res" as parent, while having
	// the lifecycle phase balance themeselves.
	var bal poststorage.Slicer
	{
		bal = tre.ObjectParent(res).ObjectLifecycle(objectlabel.LifecycleBalance)
	}

	if len(bal) == 0 {
		return nil, nil
	}

	if len(bal) == 1 {
		return bal[0], nil
	}

	return nil, tracer.Maskf(runtime.ExecutionFailedError, "too many balances for parent %s", res)
}

func ensInt(str string) (uint64, error) {
	num, err := strconv.Atoi(zerStr(str))
	if err != nil {
		return 0, tracer.Mask(err)
	}

	return uint64(num), nil
}

func zerStr(str string) string {
	if str == "" {
		str = "0"
	}

	return str
}
