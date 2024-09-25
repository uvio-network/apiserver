package postreconciler

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
	"github.com/xh3b4sd/tracer"
)

const (
	Agreement    = 0
	Disagreement = 1
	Minimum      = 2
	Creator      = 3
)

func (r *Redis) UpdateHash(pos []*poststorage.Object, hsh []string) ([]*poststorage.Object, error) {
	for i := range pos {
		if pos[i].Kind != "claim" {
			return nil, tracer.Maskf(ClaimUpdateKindError, "%s=%s", pos[i].ID, pos[i].Kind)
		}

		if !pos[i].Lifecycle.Pending() {
			return nil, tracer.Maskf(ClaimUpdateHashError, "%s=%s", pos[i].ID, pos[i].Lifecycle.Hash)
		}

		{
			pos[i].Lifecycle.Hash = append(pos[i].Lifecycle.Hash, hsh[i])
		}
	}

	return pos, nil
}

func (r *Redis) UpdateMeta(pos []*poststorage.Object, met []string) ([]*poststorage.Object, error) {
	for i := range pos {
		pos[i].Meta = met[i]
	}

	return pos, nil
}

func (r *Redis) UpdateBalance(bal *poststorage.Object, hsh []common.Hash) error {
	var err error

	for _, x := range hsh {
		bal.Lifecycle.Hash = append(bal.Lifecycle.Hash, x.Hex())
	}

	{
		err = r.sto.Post().UpdatePost([]*poststorage.Object{bal})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *Redis) UpdateResolve(res *poststorage.Object, hsh common.Hash, all []common.Address) error {
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

func (r *Redis) UpdateVotes(vot []*votestorage.Object) ([]*poststorage.Object, error) {
	var err error
	var pos []*poststorage.Object

	var sli votestorage.Slicer
	{
		sli = votestorage.Slicer(vot)
	}

	// Ensure that we do only update vote summaries for vote objects that have
	// already been confirmed onchain. Any pending votes will be ignored.
	{
		{
			sli = sli.LifecycleConfirmed()
		}

		if len(sli) == 0 {
			return nil, nil
		}
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

		if len(cla) != len(vot) {
			return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(cla), len(vot))
		}

		for i := range vot {
			pos = append(pos, updateSummary(cla[i], vot[i]))
		}
	}

	// Here we update the vote summary for the comments made on the claims that
	// votes have been cast on. The comments we update here are defined by the
	// owner of the given votes and their referenced parent claim.
	{
		var com poststorage.Slicer
		{
			com, err = r.sto.Post().SearchOwnerComment(sli.Owner(), sli.Claim())
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		for _, x := range vot {
			var y *poststorage.Object
			{
				y = com.IDClaim(x.Claim)
			}

			if y != nil {
				pos = append(pos, updateSummary(y, x))
			}
		}
	}

	return pos, nil
}

func (r *Redis) searchAddress(add []common.Address) (map[string]string, error) {
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

	return fmt.Sprintf("# Market Resolution\n\n%d %s been randomly selected to verify events in the real world for the proposed claim below.", num, use)
}

// updateSummary is responsible for incrementing the vote summary of the provided
// post objects according to the values provided with the given vote objects.
func updateSummary(pos *poststorage.Object, vot *votestorage.Object) *poststorage.Object {
	{
		// Whether the post object is of kind "claim" or kind "comment", increment
		// the "agreement" part of the vote summary by every vote value for which
		// its option is true.
		//
		//     x, 0, 0, 0
		//     x, 0
		//
		if vot.Option {
			pos.Summary[Agreement] += vot.Value
		}

		// Whether the post object is of kind "claim" or kind "comment", increment
		// the "disagreement" part of the vote summary by every vote value for
		// which its option is false.
		//
		//     0, x, 0, 0
		//     0, x
		//
		if !vot.Option {
			pos.Summary[Disagreement] += vot.Value
		}
	}

	if pos.Kind == "claim" {
		// For all claims, set the "minimum" part of the vote summary to the vote
		// value, if its value is still zero. If the "minimum" part is still zero,
		// it means this claim is updated with its first vote, which is the vote
		// of the claim creator.
		//
		//     0, 0, x, 0
		//
		if pos.Summary[Minimum] == 0 {
			pos.Summary[Minimum] += vot.Value
		}

		// For all claims, increment the "creator" part of the vote summary by the
		// vote value, if the claim owner and the vote owner are the same.
		//
		//     0, 0, 0, x
		//
		if pos.Owner == vot.Owner {
			pos.Summary[Creator] += vot.Value
		}
	}

	return pos
}
