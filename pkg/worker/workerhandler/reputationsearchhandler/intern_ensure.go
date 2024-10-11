package reputationsearchhandler

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	var dis *poststorage.Object
	{
		dis, err = h.searchClaims(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cla claimscontract.Interface
	{
		cla = h.con.Claims(dis.Contract)
	}

	var ind [8]*big.Int

	if cla.Version() == "v0.4.0" {
		lis, err := cla.SearchIndicesDeprecated(dis.ID)
		if err != nil {
			return tracer.Mask(err)
		}

		copy(ind[:], lis)
	}

	if cla.Version() == "v0.5.0" {
		ind, err = cla.SearchIndices(dis.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if tas.Sync == nil {
		tas.Sync = &task.Sync{}
	}

	var cur *big.Int
	{
		cur, err = ensBig(tas.Sync.Get(task.Paging))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var lef *big.Int
	var rig *big.Int
	var end bool
	{
		lef, rig, end, err = sample.Paging(ind, cur)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var sta []*walletstorage.Object
	{
		sta, err = h.searchWallet(cla, dis.ID, lef, rig)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Look for 10 of the highest reputation users that we know about in the
	// system.
	var rep []*userstorage.Object
	{
		rep, err = h.sto.User().SearchReputation(0, 10)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var sam *big.Int
	var wal *walletstorage.Object
	var use *userstorage.Object
	{
		sam, wal, use = h.searchSample(lef, sta, rep)
	}

	// If we just processed the last batch, or if we just found a privileged
	// voter, then emit a task respectively. If a voter exists, it will be
	// appended to the task meta data.
	if end || sam != nil {
		var blc uint64
		{
			blc, err = h.con.Client().BlockNumber(context.Background())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.emi.Claim().ResolveCreate(blc, dis.ID, ensSam(sam)...)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if sam != nil {
			h.log.Log(
				context.Background(),
				"level", "info",
				"message", "emitted task to create privileged resolve",
				"dispute", dis.ID.String(),
				"index", sam.String(),
				"address", wal.Address,
				"user", use.ID.String(),
			)
		}

		// Reset the paging pointer and all sync data in order to signal that this
		// task has completed.
		{
			tas.Sync = nil
		}

		return nil
	}

	// If no privileged voter could be found during this batch, and if there are
	// more batches to process, then update the paging pointer for the next batch
	// and wait for the next round of execution.
	{
		tas.Sync.Set(task.Paging, rig.Add(rig, big.NewInt(1)).String())
	}

	return nil
}

func (h *InternHandler) searchClaims(tas *task.Task) (*poststorage.Object, error) {
	var err error

	// The claim ID obtained here is the ID of the dispute that expired when the
	// task that we are processing right now got emitted.
	var cla objectid.ID
	{
		cla = objectid.ID(tas.Meta.Get(objectlabel.ClaimObject))
	}

	var pos poststorage.Slicer
	{
		pos, err = h.sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != 1 {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one post object for ID %s", cla)
	}

	// This is the claim with lifecycle phase "dispute" that has been expired.
	return pos[0], nil
}

func (h *InternHandler) searchSample(off *big.Int, wal []*walletstorage.Object, rep []*userstorage.Object) (*big.Int, *walletstorage.Object, *userstorage.Object) {
	dic := map[objectid.ID]*userstorage.Object{}
	for i := range rep {
		dic[rep[i].ID] = rep[i]
	}

	for i, x := range wal {
		var use *userstorage.Object
		var exi bool
		{
			use, exi = dic[x.Owner]
		}

		if exi {
			return big.NewInt(0).Add(off, big.NewInt(int64(i))), x, use
		}
	}

	return nil, nil, nil
}

func (h *InternHandler) searchWallet(cla claimscontract.Interface, dis objectid.ID, lef *big.Int, rig *big.Int) ([]*walletstorage.Object, error) {
	var err error

	var sta []common.Address
	{
		sta, err = cla.SearchStakers(dis, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var str []string
	for _, x := range sta {
		str = append(str, x.Hex())
	}

	var wal []*walletstorage.Object
	{
		wal, err = h.sto.Wallet().SearchAddress(str)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(wal) != len(sta) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(wal), len(sta))
	}

	return wal, nil
}

func ensBig(str string) (*big.Int, error) {
	val, ok := new(big.Int).SetString(zerStr(str), 10)
	if !ok {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "cannot convert %s to *big.Int", str)
	}

	return val, nil
}

func ensSam(sam *big.Int) []string {
	if sam == nil {
		return nil
	}

	return []string{sam.String()}
}

func zerStr(str string) string {
	if str == "" {
		str = "0"
	}

	return str
}
