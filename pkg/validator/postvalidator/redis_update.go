package postvalidator

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
	{
		pos, err = r.sto.Post().SearchPost(votestorage.Slicer(vot).Claim())
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != len(vot) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(pos), len(vot))
	}

	return updateVotes(pos, vot), nil
}

// updateVotes is responsible for incrementing the vote summary of the provided
// post objects according to the values provided with the given vote objects.
func updateVotes(pos []*poststorage.Object, vot []*votestorage.Object) []*poststorage.Object {
	for i := range pos {
		{
			// Whether the post object is of kind "claim" or kind "comment", increment
			// the "agreement" part of the vote summary by every vote value for which
			// its option is true.
			//
			//     x, 0, 0, 0
			//     x, 0
			//
			if vot[i].Option {
				pos[i].Votes[Agreement] += vot[i].Value
			}

			// Whether the post object is of kind "claim" or kind "comment", increment
			// the "disagreement" part of the vote summary by every vote value for
			// which its option is false.
			//
			//     0, x, 0, 0
			//     0, x
			//
			if !vot[i].Option {
				pos[i].Votes[Disagreement] += vot[i].Value
			}
		}

		if pos[i].Kind == "claim" {
			// For all claims, set the "minimum" part of the vote summary to the vote
			// value, if its value is still zero. If the "minimum" part is still zero,
			// it means this claim is updated with its first vote, which is the vote
			// of the claim creator.
			//
			//     0, 0, x, 0
			//
			if pos[i].Votes[Minimum] == 0 {
				pos[i].Votes[Minimum] += vot[i].Value
			}

			// For all claims, increment the "creator" part of the vote summary by the
			// vote value, if the claim owner and the vote owner are the same.
			//
			//     0, 0, 0, x
			//
			if pos[i].Owner == vot[i].Owner {
				pos[i].Votes[Creator] += vot[i].Value
			}
		}
	}

	return pos
}
