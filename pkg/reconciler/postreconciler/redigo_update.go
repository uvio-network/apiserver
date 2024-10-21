package postreconciler

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/uvio-network/apiserver/pkg/summary"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) UpdateHash(pos []*poststorage.Object, hsh []string) ([]*poststorage.Object, error) {
	for i := range pos {
		if pos[i].Kind != "claim" {
			return nil, tracer.Maskf(ClaimUpdateKindError, "%s = %s", pos[i].ID, pos[i].Kind)
		}

		if !pos[i].Lifecycle.Pending() {
			return nil, tracer.Maskf(ClaimUpdateHashError, "%s = %s", pos[i].ID, pos[i].Lifecycle.Hash)
		}

		{
			pos[i].Lifecycle.Hash = append(pos[i].Lifecycle.Hash, hsh[i])
		}
	}

	return pos, nil
}

func (r *Redigo) UpdateMeta(pos []*poststorage.Object, met []string) ([]*poststorage.Object, error) {
	for i := range pos {
		pos[i].Meta = met[i]
	}

	return pos, nil
}

func (r *Redigo) UpdateBalance(bal *poststorage.Object, hsh []common.Hash, sum []float64) error {
	var err error

	for _, x := range hsh {
		bal.Lifecycle.Hash = append(bal.Lifecycle.Hash, x.Hex())
	}

	{
		bal.Summary = sum
	}

	if sum[2] == 1 {
		var dis string
		if sum[3] == 0 {
			dis = "without any disputes."
		} else {
			dis = fmt.Sprintf("after %d disputes.", int(sum[3]))
		}

		var res string
		if sum[0] == 1 {
			res = fmt.Sprintf(" The associated claim was found to be %s %s", txtTru(sum[1]), dis)
		}

		bal.Text = fmt.Sprintf("This market was finalized with %s resolution.%s", txtVal(sum[0]), res)
	}

	{
		err = r.sto.Post().UpdatePost([]*poststorage.Object{bal})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *Redigo) UpdateResolve(res *poststorage.Object, hsh common.Hash, all []common.Address) error {
	var err error

	{
		res.Text = resTxt(len(all))
	}

	if len(hsh) != 0 {
		res.Lifecycle.Hash = []string{hsh.Hex()}
	}

	{
		sam, err := r.searchAddress(all)
		if err != nil {
			return tracer.Mask(err)
		}

		if res.Samples == nil {
			res.Samples = map[string]string{}
		}

		for k, v := range sam {
			res.Samples[k] = v
		}
	}

	{
		err = r.sto.Post().UpdatePost([]*poststorage.Object{res})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *Redigo) UpdateVotes(vot []*votestorage.Object) ([]*poststorage.Object, error) {
	var err error
	var pos []*poststorage.Object

	var sli votestorage.Slicer
	{
		sli = votestorage.Slicer(vot).LifecycleConfirmed()
	}

	// Ensure that we do only update vote summaries for vote objects that have
	// already been confirmed onchain. Any pending votes will be ignored.
	if len(sli) == 0 {
		return nil, nil
	}

	// Here we update the vote summary for the claims on which the provided votes
	// have been cast. The claims we update here are defined by the parent of the
	// given votes.
	{
		var cla []*poststorage.Object
		{
			cla, err = r.sto.Post().SearchPost(sli.Claim())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		if len(cla) != len(sli) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(cla), len(sli))
		}

		for i := range sli {
			pos = append(pos, summary.Update(cla[i], sli[i]))
		}
	}

	// Here we update the vote summary for the comments made on the claims that
	// votes have been cast on. The comments we update here are defined by the
	// owner of the given votes and their referenced parent claim. In other words,
	// if I have two comments on a claim already, and I stake some more tokens on
	// that claim again, then my two comments need to get updated in order to
	// reflect the total amount of skin in the game that I have in this market.
	{
		var com poststorage.Slicer
		{
			com, err = r.sto.Post().SearchOwnerComment(sli.Owner(), sli.Claim())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// Find the comments that have the same parent as the votes that we are
		// processing here, because those votes carry the values that we need to add
		// to the comments' vote summaries.
		for _, x := range sli {
			for _, y := range com.ObjectParent(x.Claim) {
				pos = append(pos, summary.Update(y, x))
			}
		}
	}

	return pos, nil
}

func (r *Redigo) searchAddress(add []common.Address) (map[string]string, error) {
	var err error

	var str []string
	for _, x := range add {
		str = append(str, x.Hex())
	}

	var wal []*walletstorage.Object
	{
		wal, err = r.sto.Wallet().SearchAddress(str)
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

func resTxt(num int) string {
	var use string
	{
		use = "user has"
	}

	if num > 1 {
		use = "users have"
	}

	return fmt.Sprintf("%d %s been randomly selected to verify events in the real world for the proposed claim below.", num, use)
}

func txtTru(f float64) string {
	if f == 1 {
		return "**true**"
	}

	return "**false**"
}

func txtVal(f float64) string {
	if f == 1 {
		return "a **valid**"
	}

	return "an **invalid**"
}
