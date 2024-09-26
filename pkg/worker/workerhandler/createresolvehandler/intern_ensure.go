package createresolvehandler

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
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
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

	// oneWeek is the time within which selected stakers can vote on the created
	// resolves. The resolve expiries start from the associated propose expiry and
	// last until this 7 standard days later.
	oneWeek = time.Hour * 24 * 7
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var pod *poststorage.Object
	var res *poststorage.Object
	{
		pod, res, err = h.searchClaims(tas)
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

	var exp time.Time
	{
		exp = pod.Expiry.Add(oneWeek)
	}

	// If the next claim within this tree relative to the provided claim is nil,
	// then that means that we have not yet created the required resolve. And so
	// we can proceed to create the post object for the requested resolve.
	if res == nil {
		res, err = h.rec.Post().CreateResolve(pod, exp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var exi bool
	{
		exi, err = cla.ExistsResolve(pod.ID)
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

	var hsh common.Hash
	if !exi {
		var sam []*big.Int
		{
			sam = sample.BigInt(h.sam.Random(ind[0].Uint64(), ind[7].Uint64()))
		}

		var txn *types.Transaction
		{
			txn, err = cla.CreateResolve(pod.ID, sam, exp)
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

		{
			hsh = txn.Hash()
		}
	} else {
		var blc uint64
		{
			blc, err = tasInt(tas.Sync.Get(task.Paging))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		hsh, err = cla.ResolveCreated(blc, pod.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var tru []common.Address
	var fls []common.Address
	{
		tru, fls, err = h.searchSamples(cla, pod.ID, ind)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if res.Lifecycle.Pending() {
		err = h.rec.Post().UpdateResolve(res, hsh, append(tru, fls...))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Once the new post object got updated with the transaction hash and the
	// sampled addresses we can remove the paging pointer and all sync state from
	// the task so that it can be deleted eventually by the rescue engine.
	{
		tas.Sync = nil
	}

	// TODO we need to notify the voters somehow so that they know they have to
	// verify events in the real world

	return nil
}

func (h *InternHandler) searchClaims(tas *task.Task) (*poststorage.Object, *poststorage.Object, error) {
	var err error

	// The claim ID obtained here is the ID of the propose or dispute that expired
	// when the task that we are processing right now got emitted.
	var cla objectid.ID
	{
		cla = objectid.ID(tas.Meta.Get(objectlabel.ClaimObject))
	}

	var pos poststorage.Slicer
	{
		pos, err = h.sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	// This is the claim with lifecycle phase "propose" or "dispute" that has been
	// expired.
	var pod *poststorage.Object
	{
		pod = pos.IDClaim(cla)
	}

	var tre poststorage.Slicer
	{
		tre, err = h.sto.Post().SearchTree([]objectid.ID{pod.Tree})
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	var res *poststorage.Object
	{
		res = tre.NextClaim(pod.ID)
	}

	return pod, res, nil
}

func (h *InternHandler) searchSamples(cla claimscontract.Interface, pod objectid.ID, ind []*big.Int) ([]common.Address, []common.Address, error) {
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
		tru, err = cla.SearchSamples(pod, ind[1], ind[2])
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	var fls []common.Address
	{
		fls, err = cla.SearchSamples(pod, ind[5], ind[6])
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	return tru, fls, nil
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
