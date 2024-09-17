package uvxminthandler

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	// maxWait is the maximum amount of time to wait for transactions per task
	// execution. We need to limit the task execution time because every worker
	// has only 30 seconds to process a task by default before ownership of the
	// provided task will be revoked. Therefore defining maxWait under 30 seconds
	// gives us enough time to gracefully process the task given to us.
	maxWait = 20 * time.Second
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var use objectid.ID
	{
		use = objectid.ID(tas.Meta.Get(objectlabel.UserObject))
	}

	var wal walletstorage.Slicer
	{
		wal, err = h.sto.Wallet().SearchOwner([]objectid.ID{use})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var con walletstorage.Slicer
	{

		con = wal.ObjectActive(true)
	}

	// Anyone not having a wallet cannot receive minted UVX tokens. We are not
	// returning an error here in case no wallet can be found, because that would
	// produce a deadlock scenario in which tasks would be repeatedly called
	// forever, only to end up in the same situation. Instead we are using the the
	// task's sync paging pointer as a delayed retry signal.
	if len(con) == 0 {
		var cur string
		{
			cur = tas.Sync.Get(task.Paging)
		}

		var des string
		{
			des = createPaging(cur)
		}

		if tas.Sync == nil {
			tas.Sync = &task.Sync{}
		}

		{
			tas.Sync.Set(task.Paging, des)
		}

		{
			time.Sleep(1 * time.Second)
		}

		return nil
	}

	// It may happen that no user wallet could be found intermittently. In such a
	// case we may have set a paging pointer in order to retry the task. If we
	// retried and then found a wallet for the user that we are asked to mint
	// tokens for, then we need to make sure that we only do that once. So just to
	// make sure that we are not minting tokens again once we got to this point
	// here, we are ensuring that the paging pointer is removed, by simply setting
	// the task's sync map to nil, which erases all sync state entirely. That
	// tells the rescue engine to finally delete the task, unless we return an
	// error below.
	{
		tas.Sync = nil
	}

	var add string
	{
		add = con[0].Address
	}

	var txn *types.Transaction
	{
		txn, err = h.con.UVX().Mint(add, 100)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// In case we are unable to wait for the minting transaction to be included in
	// a block onchain, we are not returning an error, since this would
	// potentially cause a continuous mint for the given user. The safest thing to
	// do is to stop the process here and allow a human to mint tokens to the
	// address if the given address has not received any tokens yet. Further, if
	// we end up logging an error here, the root cause for this error must be
	// investigated by a human. The spectrum of potential errors can range from
	// network hickups to funding issues, where the underlying signer may ran out
	// of ETH used to pay for gas.
	{
		err = h.con.Wait(txn, maxWait)
		if err != nil {
			h.log.Log(
				context.Background(),
				"level", "error",
				"message", err.Error(),
				"description", "Waiting for the mint transaction failed. The root cause for this failure needs to be investigated. The given address may not have received UVX tokens as requested.",
				"address", add,
				"amount", "100",
				"transaction", txn.Hash().Hex(),
				"stack", tracer.Stack(err),
			)
		}
	}

	return nil
}

func createPaging(cur string) string {
	if cur == "1" {
		return "2"
	}

	if cur == "2" {
		return "0"
	}

	return "1"
}
