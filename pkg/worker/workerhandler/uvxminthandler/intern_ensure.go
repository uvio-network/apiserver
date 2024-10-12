package uvxminthandler

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/contract/uvxcontract"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
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

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
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

	var bal *big.Int
	{
		bal, err = uvx.Balance(add)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Prevent users from getting tokens multiple times. If we are processing a
	// task for a user who has already a token balance, then we have nothing else
	// to do here.
	if bal.Uint64() > 0 {
		return nil
	}

	var txn *types.Transaction
	{
		txn, err = uvx.Mint(add, 100)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// In case the minting transaction fails, we are going to retry this task.
	{
		err = h.con.Wait(txn, maxWait)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
