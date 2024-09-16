package createresolvehandler

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/logger/meta"
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

const (
	oneWeek = time.Hour * 24 * 7
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var cyc int
	{
		cyc, err = strconv.Atoi(tas.Meta.Get(task.Cycles))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if cyc > 10 {
		h.log.Log(
			logCtx(tas),
			"level", "error",
			"message", "task cycle limit error",
			"description", "Creating the resolve onchain failed. The root cause for this failure needs to be investigated. The given propose may still not have a resolve as requested.",
		)

		return nil
	}

	// The claim ID obtained here is the ID of the propose that expired when the
	// task that we are processing right now got emitted.
	var cla objectid.ID
	{
		cla = objectid.ID(tas.Meta.Get(objectlabel.ClaimObject))
	}

	var tre poststorage.Slicer
	{
		tre, err = h.sto.Post().SearchTree([]objectid.ID{cla})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var pro *poststorage.Object
	var res *poststorage.Object
	{
		pro = tre.IDClaim(cla)
		res = tre.NextClaim(cla)
	}

	var exp time.Time
	{
		exp = time.Now().Add(oneWeek)
	}

	// If the next claim within this tree relative to the provided propose is nil,
	// then that means that we have not yet created the required resolve. And so
	// we can proceed to create the post object for the requested resolve.
	if res == nil {
		err = h.createObject(pro, exp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var exi bool
	{
		exi, err = h.con.Claims().ExistsResolve(pro.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var ind []*big.Int
	{
		ind, err = h.con.Claims().SearchIndices(pro.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var txn *types.Transaction
	if !exi {
		var sam []*big.Int
		{
			sam = sample.BigInt(h.sam.Random(ind[0].Uint64(), ind[7].Uint64()))
		}

		{
			txn, err = h.con.Claims().CreateResolve(pro.ID, sam, exp)
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
	}

	var tru []common.Address
	var fls []common.Address
	{
		tru, fls, err = h.searchSamples(pro.ID, ind)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if res.Lifecycle.Pending() {
		err = h.updateObject(res, txn, tru, fls)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// TODO we need to notify the voters somehow so that they know they have to
	// verify events in the real world

	return nil
}

func (h *InternHandler) createObject(pro *poststorage.Object, exp time.Time) error {
	var err error

	var res *poststorage.Object
	{
		res = &poststorage.Object{
			Chain:  pro.Chain,
			Expiry: exp,
			Kind:   "claim",
			Labels: pro.Labels,
			Lifecycle: objectfield.Lifecycle{
				Data: objectlabel.LifecycleResolve,
			},
			Owner:  objectid.System(),
			Parent: pro.ID,
			Text:   "# Market Resolution\n\n The random truth sampling process has begun and is waiting for onchain confirmation.",
		}
	}

	var out []*poststorage.Object
	{
		out, err = h.rec.Post().CreatePost([]*poststorage.Object{res})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.sto.Post().CreatePost(out)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) searchAddress(add []common.Address) (map[string]string, error) {
	var err error

	var str []string
	for _, x := range add {
		str = append(str, x.Hex())
	}

	var wal []*walletstorage.Object
	{
		wal, err = h.sto.Wallet().SearchAddress(str)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(wal) != len(add) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(wal), len(add))
	}

	dic := map[string]string{}

	for i := range add {
		dic[str[i]] = string(wal[i].Owner)
	}

	return dic, nil
}

func (h *InternHandler) searchSamples(cla objectid.ID, ind []*big.Int) ([]common.Address, []common.Address, error) {
	var err error

	if len(ind) != 8 {
		return nil, nil, tracer.Maskf(runtime.ExecutionFailedError, "onchain indices invalid")
	}

	// Searching for the sampled addresses works according to the indices [1, 2]
	// and [5, 6] according to the Claims contract documentation linked below.
	//
	//     https://github.com/uvio-network/contracts/blob/v0.2.0/contracts/Claims.sol#L1721-L1728
	//

	var tru []common.Address
	{
		tru, err = h.con.Claims().SearchSamples(cla, ind[1], ind[2])
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	var fls []common.Address
	{
		fls, err = h.con.Claims().SearchSamples(cla, ind[5], ind[6])
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	return tru, fls, nil
}

func (h *InternHandler) updateObject(res *poststorage.Object, txn *types.Transaction, tru []common.Address, fls []common.Address) error {
	var err error

	var num int
	{
		num = len(tru) + len(fls)
	}

	{
		res.Text = resTxt(num)
	}

	if txn != nil {
		res.Lifecycle.Hash = []string{txn.Hash().Hex()}
	}

	{
		var tmp map[string]string
		if len(tru) != 0 {
			tmp, err = h.searchAddress(tru)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var fmp map[string]string
		if len(fls) != 0 {
			fmp, err = h.searchAddress(fls)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			res.Samples = map[string]map[string]string{
				"true":  tmp,
				"false": fmp,
			}
		}
	}

	{
		err = h.sto.Post().UpdatePost([]*poststorage.Object{res})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func logCtx(tas *task.Task) context.Context {
	ctx := context.Background()

	for k, v := range *tas.Meta {
		ctx = meta.Add(ctx, k, v)
	}

	return ctx
}

func resTxt(num int) string {
	var use string
	{
		use = "user"
	}

	if num > 1 {
		use += "s"
	}

	return fmt.Sprintf("# Market Resolution\n\n %d %s have been randomly selected to verify events in the real world for the proposed claim below.", num, use)
}
