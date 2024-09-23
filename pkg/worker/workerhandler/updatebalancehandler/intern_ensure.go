package updatebalancehandler

import (
	"context"
	"math/big"
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
	"github.com/xh3b4sd/logger/meta"
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

	{
		var cyc uint64
		{
			cyc, err = tasInt(tas.Meta.Get(task.Cycles))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if cyc > 10 {
			h.log.Log(
				logCtx(tas),
				"level", "error",
				"message", "task cycle limit error",
				"description", "Settling the propose onchain failed. The root cause for this failure needs to be investigated. The given propose may still not be settled as requested.",
			)

			return nil
		}
	}

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

	{
		var blc uint64
		{
			blc, err = cla.Client().BlockNumber(context.Background())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if tas.Sync == nil {
			tas.Sync = &task.Sync{}
		}

		{
			tas.Sync.Set(task.Paging, strconv.FormatInt(int64(blc), 10))
		}
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

	var hsh []common.Hash
	{
		var blc uint64
		{
			blc, err = tasInt(tas.Sync.Get(task.Paging))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		hsh, err = cla.BalanceUpdated(blc, uint64(pod.ID.Int()))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var ind []*big.Int
	{
		ind, err = cla.SearchIndices(pod.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var tru int64
	var fls int64
	{
		tru, fls, err = cla.SearchVotes(pod.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var sum []float64
	{
		sum = []float64{
			float64(tru),
			float64(fls),
			float64(ind[0].Int64()),
			float64(ind[7].Int64()),
		}
	}

	if bal.Lifecycle.Pending() {
		err = h.rec.Post().UpdateBalance(bal, hsh, sum)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Once the new post object got updated with all the associated transaction
	// hashes, we can remove the paging pointer and all sync state from the task
	// so that it can be deleted eventually by the rescue engine.
	{
		tas.Sync = nil
	}

	// TODO we need to notify the winners somehow so that they know they have won
	// something

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

	// This is the claim with lifecycle phase resolve that has been expired.
	var res *poststorage.Object
	{
		res = pos.IDClaim(cla)
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
		bal = tre.NextClaim(res.ID)
	}

	// If the balance object exists and it is not of lifecycle phase "balance",
	// then something went horribly wrong across the claim reconciliation. The
	// only possible lifecycle phase imaginable would be "dispute", if something
	// went wrong with deferring the resolve expiry in the corresponding expiry
	// queue.
	if bal != nil && !bal.Lifecycle.Is(objectlabel.LifecycleBalance) {
		return nil, nil, nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	return pod, res, bal, nil
}

func logCtx(tas *task.Task) context.Context {
	ctx := context.Background()

	for k, v := range *tas.Meta {
		ctx = meta.Add(ctx, k, v)
	}

	return ctx
}

func tasInt(str string) (uint64, error) {
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
