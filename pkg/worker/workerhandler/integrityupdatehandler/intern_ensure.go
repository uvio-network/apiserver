package integrityupdatehandler

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
	"golang.org/x/mod/semver"
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

	// We can only update reputation metrics if the claims that we are processing
	// have been facilitated by Claims contracts of at least version "v0.5.0".
	// Contracts equal to that version or greater than that version yield 0 or +1
	// respectively when calling semver.Compare.
	if semver.Compare("v0.5.0", cla.Version()) == -1 {
		return nil
	}

	//
	//     market                                   compute units
	//
	//     searchResults(true) (bool, ...)          10
	//     searchIndices(true) (uint256, ...)       10
	//
	//     competence                               compute units
	//
	//     searchStakers(true) []Address            10
	//     searchStakers(false) []Address           10
	//     searchHistory(true) []uint256            10
	//     searchHistory(false) []uint256           10
	//
	//     integrity                                compute units
	//
	//     searchVoters(true) []Address             10
	//     searchVoters(false) []Address            10
	//     searchSamples(true) []uint256            10
	//     searchSamples(false) []uint256           10
	//
	//     total                                    100
	//

	var val bool
	var sid bool
	var fin bool
	{
		val, sid, fin, err = cla.SearchResults(pod.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if !fin {
		return tracer.Maskf(runtime.ExecutionFailedError, "claim tree %s not finalized onchain", pod.Tree)
	}

	var ind [8]*big.Int
	{
		ind, err = cla.SearchIndices(pod.ID)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var vot []common.Address
	{
		vot, err = h.searchVoters(cla, pod.ID, ind)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var sam []uint8
	{
		sam, err = h.searchSamples(cla, pod.ID, ind)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if len(vot) != len(sam) {
		return tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(vot), len(sam))
	}

	var use map[string]*userstorage.Object
	{
		use, err = h.searchUsers(h.sto, res)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if len(vot) != len(use) {
		return tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(vot), len(use))
	}

	// For every key, which is the address of a voter, we update the respective
	// user object according to the found integrity metrics.

	for i := range vot {
		var v uint8
		{
			v = sam[i]
		}

		var a string
		{
			a = vot[i].Hex()
		}

		var u *userstorage.Object
		{
			u = use[a]
		}

		// If the user denied to vote for whatever reason, then increment the
		// abstained value.
		if v == 2 {
			u.Summary[userstorage.Abstained]++
		}

		if val {
			if sid {
				// If the market settled with a valid resolution, and if the community
				// consensus was true, and if the user voted true, then increment the
				// honest value.
				if v == 1 {
					u.Summary[userstorage.Honest]++
				}
			} else {
				// If the market settled with a valid resolution, and if the community
				// consensus was false, and if the user voted false, then increment the
				// honest value.
				if v == 0 {
					u.Summary[userstorage.Honest]++
				}
			}
		} else {
			// If the market settled with an invalid resolution, then increment the
			// dishonest value.
			u.Summary[userstorage.Dishonest]++
		}

		{
			err = h.sto.User().UpdateUser([]*userstorage.Object{u})
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (h *InternHandler) searchClaims(tas *task.Task) (*poststorage.Object, *poststorage.Object, error) {
	var err error

	// The claim ID obtained here is the ID of the balance that finalized when the
	// task that we are processing right now got emitted.
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

	if len(pos) != 1 {
		return nil, nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one post object for ID %s", cla)
	}

	// This is the claim with lifecycle phase balance that has been finalized.
	var bal *poststorage.Object
	{
		bal = pos[0]
	}

	var tre poststorage.Slicer
	{
		tre, err = h.sto.Post().SearchTree([]objectid.ID{bal.Tree})
		if err != nil {
			return nil, nil, tracer.Mask(err)
		}
	}

	// This claim here is the resolve of the balance that got finalized.
	var res *poststorage.Object
	{
		res = tre.IDObject(bal.Parent)
	}

	if res == nil {
		return nil, nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one resolve for ID %s", bal.Parent)
	}

	// This claim here is either a propose or a dispute in first or second
	// instance. In either case, those are the claims that we are tasked to
	// settle.
	var pod *poststorage.Object
	{
		pod = tre.IDObject(res.Parent)
	}

	if pod == nil {
		return nil, nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one propose for ID %s", res.Parent)
	}

	return pod, res, nil
}

func (h *InternHandler) searchSamples(cla claimscontract.Interface, pod objectid.ID, ind [8]*big.Int) ([]uint8, error) {
	var err error

	// Searching for the voter outcomes works using the indices [1, 2] and [5, 6]
	// according to the Claims contract documentation linked below.
	//
	//     https://github.com/uvio-network/contracts/blob/v0.5.0/contracts/Claims.sol#L1826-L1882
	//

	var yay []uint8
	{
		yay, err = cla.SearchSamples(pod, ind[1], ind[2])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var nah []uint8
	{
		nah, err = cla.SearchSamples(pod, ind[5], ind[6])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return append(yay, nah...), nil
}

func (h *InternHandler) searchUsers(sto storage.Interface, res *poststorage.Object) (map[string]*userstorage.Object, error) {
	var err error

	var add []string
	var uid []objectid.ID
	for k, v := range res.Samples {
		add = append(add, k)
		uid = append(uid, objectid.ID(v))
	}

	var use []*userstorage.Object
	{
		use, err = sto.User().SearchUser(uid)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(add) != len(use) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(add), len(use))
	}

	dic := map[string]*userstorage.Object{}
	for i := range add {
		dic[add[i]] = use[i]
	}

	return dic, nil
}

func (h *InternHandler) searchVoters(cla claimscontract.Interface, pod objectid.ID, ind [8]*big.Int) ([]common.Address, error) {
	var err error

	// Searching for the voter addresses works using the indices [1, 2] and [5, 6]
	// according to the Claims contract documentation linked below.
	//
	//     https://github.com/uvio-network/contracts/blob/v0.5.0/contracts/Claims.sol#L1826-L1882
	//

	var yay []common.Address
	{
		yay, err = cla.SearchVoters(pod, ind[1], ind[2])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var nah []common.Address
	{
		nah, err = cla.SearchVoters(pod, ind[5], ind[6])
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return append(yay, nah...), nil
}
