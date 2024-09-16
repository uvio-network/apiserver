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
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
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
			logctx(tas),
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

	var exp time.Time
	{
		exp = time.Now().Add(oneWeek)
	}

	var exi bool
	{
		exi, err = h.con.Claims().ExistsResolve(cla)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var lef uint64
	var rig uint64
	{
		lef, rig, err = h.con.Claims().SearchIndices(cla)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var txn *types.Transaction
	if !exi {
		var ind []*big.Int
		{
			ind = sample.BigInt(h.sam.Random(lef, rig))
		}

		{
			txn, err = h.con.Claims().CreateResolve(cla, ind, exp)
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

	var add []common.Address
	{
		add, err = h.con.Claims().SearchSamples(cla, lef, rig)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// TODO we need to notify the voters somehow so that they know they have to
	// verify events in the real world

	{
		err = h.ensureObject(cla, txn, add, exp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) ensureObject(cla objectid.ID, txn *types.Transaction, add []common.Address, exp time.Time) error {
	var err error

	var tre poststorage.Slicer
	{
		tre, err = h.sto.Post().SearchTree([]objectid.ID{cla})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// If the next claim within this tree relative to the provided propose is not
	// nil, then that means that we have already created the required resolve. And
	// so we can stop processing here already.
	if tre.NextClaim(cla) != nil {
		return nil
	}

	var has []string
	if txn != nil {
		has = []string{txn.Hash().Hex()}
	}

	var pro *poststorage.Object
	{
		pro = tre.IDClaim(cla)
	}

	var res *poststorage.Object
	{
		res = &poststorage.Object{
			Chain:  pro.Chain,
			Expiry: exp,
			Kind:   "claim",
			Labels: pro.Labels,
			Lifecycle: objectfield.Lifecycle{
				Data: objectlabel.LifecycleResolve,
				Hash: has,
			},
			Owner:  objectid.System(),
			Parent: pro.ID,
			Text:   resTxt(add),
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

func resTxt(add []common.Address) string {
	var use string
	{
		use = "user"
	}

	if len(add) > 1 {
		use += "s"
	}

	return fmt.Sprintf("# Market Resolution\n\n %d %s have been randomly selected to verify events in the real world for the proposed claim below.", len(add), use)
}

func logctx(tas *task.Task) context.Context {
	ctx := context.Background()

	for k, v := range *tas.Meta {
		ctx = meta.Add(ctx, k, v)
	}

	return ctx
}
