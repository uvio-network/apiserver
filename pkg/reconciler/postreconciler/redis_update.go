package postreconciler

import (
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

const (
	Agreement    = 0
	Disagreement = 1
	Minimum      = 2
	Creator      = 3
)

func (r *Redis) UpdateVotes(vot []*votestorage.Object) ([]*poststorage.Object, error) {
	var err error
	var pos []*poststorage.Object

	var sli votestorage.Slicer
	{
		sli = votestorage.Slicer(vot)
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
			pos = append(pos, updateVotes(cla[i], vot[i]))
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
			for _, y := range com.ObjectID(x.Claim) {
				pos = append(pos, updateVotes(y, x))
			}
		}
	}

	return pos, nil
}

// updateVotes is responsible for incrementing the vote summary of the provided
// post objects according to the values provided with the given vote objects.
func updateVotes(pos *poststorage.Object, vot *votestorage.Object) *poststorage.Object {
	{
		// Whether the post object is of kind "claim" or kind "comment", increment
		// the "agreement" part of the vote summary by every vote value for which
		// its option is true.
		//
		//     x, 0, 0, 0
		//     x, 0
		//
		if vot.Option {
			pos.Votes[Agreement] += vot.Value
		}

		// Whether the post object is of kind "claim" or kind "comment", increment
		// the "disagreement" part of the vote summary by every vote value for
		// which its option is false.
		//
		//     0, x, 0, 0
		//     0, x
		//
		if !vot.Option {
			pos.Votes[Disagreement] += vot.Value
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
		if pos.Votes[Minimum] == 0 {
			pos.Votes[Minimum] += vot.Value
		}

		// For all claims, increment the "creator" part of the vote summary by the
		// vote value, if the claim owner and the vote owner are the same.
		//
		//     0, 0, 0, x
		//
		if pos.Owner == vot.Owner {
			pos.Votes[Creator] += vot.Value
		}
	}

	return pos
}
