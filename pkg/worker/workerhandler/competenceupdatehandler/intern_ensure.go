package competenceupdatehandler

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
	"golang.org/x/mod/semver"
)

func (h *InternHandler) Ensure(tas *task.Task) error {
	var err error

	var pod *poststorage.Object
	{
		pod, err = h.searchClaims(tas)
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
	if semver.Compare(cla.Version(), "v0.5.0") == -1 {
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

	// After everything has been reconciled, there still may have been some error
	// preventing the task to be cleaned up properly. In this case we may receive
	// a paging pointer that results in the end of the entire iteration cycle
	// without providing any further iteration boundaries. So if this is the end,
	// and if lef and rig are empty, then stop processing the task, because we
	// have already done all the work.
	if lef == nil && rig == nil && end {
		{
			tas.Sync = nil
		}

		return nil
	}

	var sta []common.Address
	{
		sta, err = cla.SearchStakers(pod.ID, lef, rig)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var his []*big.Int
	{
		his, err = cla.SearchHistory(pod.ID, lef, rig)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// For every key, which is the address of a staker, we want to update the
	// respective user object according to the competence metrics that we are
	// going to calculate below.
	var use map[string]*userstorage.Object
	{
		use, err = h.searchUsers(sta)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if len(his) != len(use)*5 {
		return tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(his), len(use)*5)
	}

	var not []*notestorage.Object
	for i := range sta {
		var k int
		var u *userstorage.Object
		{
			k, u = updUse(i, val, sid, sta, his, use)
		}

		{
			err = h.sto.User().UpdateUser([]*userstorage.Object{u})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			not = append(not, &notestorage.Object{
				Kind:     notKin(k),
				Resource: u.ID,
				Owner:    u.ID,
			})
		}

		// After each iteration, increment the current paging pointer to track our
		// progress.
		{
			fmt.Printf("pag bef %#v\n", tas.Sync.Get(task.Paging))
			tas.Sync.Set(task.Paging, lef.Add(lef, big.NewInt(1)).String())
			fmt.Printf("pag aft %#v\n", tas.Sync.Get(task.Paging))
		}
	}

	// Once all user objects got updated for this batch, we update the reputation
	// index and ensure that the now lowest reputation users get removed from the
	// index at the end of processing all relevant batches for this cycle. Note
	// that in theory we can miss to update reputation metrics for a few users of
	// any given batch, if the updating of user objects above breaks unexpectedly
	// in the middle of the loop. We rely on eventual consistency to get those
	// "lost" users updated once another claim is being resolved in which those
	// "lost" users participated in as well. The aspect of eventual consistency is
	// good enough for us now in the offchain setting because only high reputation
	// users are processed here to begin with. And those high reputation users do
	// either have a high reputation already, or will gain it eventually anyway.
	// In this particular case we prefer the benefit of updating Redis only once
	// per batch instead of per user in the loop above.
	{
		err = h.sto.User().UpdateReputation(useLis(use))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = h.sto.Note().CreateNote(not)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if end {
		{
			err = h.sto.User().DeleteReputation()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			tas.Sync = nil
		}
	}

	return nil
}

func (h *InternHandler) searchClaims(tas *task.Task) (*poststorage.Object, error) {
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
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != 1 {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one post object for ID %s", cla)
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
			return nil, tracer.Mask(err)
		}
	}

	// This claim here is the resolve of the balance that got finalized.
	var res *poststorage.Object
	{
		res = tre.IDObject(bal.Parent)
	}

	if res == nil {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one resolve for ID %s", bal.Parent)
	}

	// This claim here is either a propose or a dispute in first or second
	// instance. In either case, those are the claims that we are tasked to
	// settle.
	var pod *poststorage.Object
	{
		pod = tre.IDObject(res.Parent)
	}

	if pod == nil {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one propose for ID %s", res.Parent)
	}

	return pod, nil
}

func (h *InternHandler) searchUsers(sta []common.Address) (map[string]*userstorage.Object, error) {
	var err error

	var add []string
	for _, x := range sta {
		add = append(add, x.Hex())
	}

	var sli walletstorage.Slicer
	{
		sli, err = h.sto.Wallet().SearchAddress(add)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var use []*userstorage.Object
	{
		use, err = h.sto.User().SearchUser(sli.Owner())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	dic := map[string]*userstorage.Object{}
	for i := range add {
		dic[add[i]] = use[i]
	}

	if len(sta) != len(dic) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(sta), len(dic))
	}

	return dic, nil
}

func bigEql(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func bigGrt(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == +1
}

func bigNot(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) != 0
}

func ensBig(str string) (*big.Int, error) {
	val, ok := new(big.Int).SetString(zerStr(str), 10)
	if !ok {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "cannot convert %s to *big.Int", str)
	}

	return val, nil
}

func notKin(k int) string {
	if k == userstorage.Right {
		return "userRight"
	}

	if k == userstorage.Wrong {
		return "userWrong"
	}

	return ""
}

func updUse(i int, val bool, sid bool, sta []common.Address, his []*big.Int, use map[string]*userstorage.Object) (int, *userstorage.Object) {
	var zer *big.Int
	{
		zer = big.NewInt(0)
	}

	//
	//     [        before        |         after        |   ]
	//
	//     "agreement,disagreement,agreement,disagreement,fee"
	//
	var b []*big.Int
	{
		b = his[i*5 : (i+1)*5]
	}

	var a string
	{
		a = sta[i].Hex()
	}

	var k int

	var u *userstorage.Object
	{
		u = use[a]
	}

	if val {
		if sid {
			// If the market settled with a valid resolution, and if the community
			// consensus was true, and if the user staked true while not staking
			// false, and if the staked true balance after resolution is higher or at
			// least equal to the staked true balance before resolution, then
			// increment the right value.
			if bigNot(b[0], zer) && bigEql(b[1], zer) && (bigEql(b[2], b[0]) || bigGrt(b[2], b[0])) {
				k = userstorage.Right
				u.Summary[userstorage.Right]++
			} else {
				k = userstorage.Wrong
				u.Summary[userstorage.Wrong]++
			}
		} else {
			// If the market settled with a valid resolution, and if the community
			// consensus was false, and if the user staked false while not staking
			// true, and if the staked false balance after resolution is higher or at
			// least equal to the staked false balance before resolution, then
			// increment the honest value.
			if bigNot(b[1], zer) && bigEql(b[0], zer) && (bigEql(b[3], b[1]) || bigGrt(b[3], b[1])) {
				k = userstorage.Right
				u.Summary[userstorage.Right]++
			} else {
				k = userstorage.Wrong
				u.Summary[userstorage.Wrong]++
			}
		}
	}

	return k, u
}

func useLis(use map[string]*userstorage.Object) []*userstorage.Object {
	var lis []*userstorage.Object

	for _, v := range use {
		lis = append(lis, v)
	}

	return lis
}

func zerStr(str string) string {
	if str == "" {
		str = "0"
	}

	return str
}
