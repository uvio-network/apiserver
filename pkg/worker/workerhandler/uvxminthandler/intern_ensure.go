package uvxminthandler

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
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

		con = wal.ObjectActive(true).ObjectKind("contract")
	}

	// TODO this is a deadlock, we have to ensure every user has 1 valid wallet at
	// all times and we need some alerting to understand the broken state in here.
	if len(con) == 0 {
		return tracer.Mask(runtime.ExecutionFailedError)
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

	{
		err = h.con.Wait(txn)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
