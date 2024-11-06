package uvxminthandler

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/contract/uvxcontract"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/objectid"
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

func (h *InternHandler) Ensure(tas *task.Task) error {
	var err error

	var use objectid.ID
	{
		use = objectid.ID(tas.Meta.Get(objectlabel.UserObject))
	}

	var wal []*walletstorage.Object
	{
		wal, err = h.sto.Wallet().SearchOwner([]objectid.ID{use})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Anyone not having a wallet cannot receive minted UVX tokens. It may happen
	// that users do not have an associated wallet during signup temporarily. In
	// those cases we return an error in order to execute this task again at a
	// later point in time.
	if len(wal) == 0 {
		return tracer.Maskf(WalletNotFoundError, "%s", use)
	}

	var uvx uvxcontract.Interface
	{
		uvx = h.con.UVX(h.uvx)
	}

	var add string
	{
		add = wal[0].Address
	}

	var txn *types.Transaction
	{
		txn, err = uvx.Mint(add, 100)
		if err != nil {
			if tas.Cron == nil {
				tas.Cron = &task.Cron{}
			}

			// Ensure the failed task is being retried after some backoff period. If
			// we have retried this task already, then we also need to make sure to
			// remove the tick+1 label in order to get another try.
			{
				tas.Cron.Set().Adefer("1 minute")
				tas.Cron.Prg().TickP1()
			}

			h.log.Log(
				context.Background(),
				"level", "warning",
				"message", err.Error(),
				"stack", tracer.Stack(err),
			)

			return nil
		}
	}

	{
		err = h.con.Wait(txn, maxWait)
		if err != nil {
			h.log.Log(
				context.Background(),
				"level", "warning",
				"message", err.Error(),
				"stack", tracer.Stack(err),
			)

			return nil
		}
	}

	var not *notestorage.Object
	{
		not = &notestorage.Object{
			Kind:    "uvxMint",
			Message: fmt.Sprintf("%d UVX tokens have been minted to your address %s. Thanks for playing!", 100, add),
			Owner:   use,
		}
	}

	{
		err = h.sto.Note().CreateNote([]*notestorage.Object{not})
		if err != nil {
			h.log.Log(
				context.Background(),
				"level", "warning",
				"message", err.Error(),
				"stack", tracer.Stack(err),
			)

			return nil
		}
	}

	return nil
}
