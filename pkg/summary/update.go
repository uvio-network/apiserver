package summary

import (
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
)

// Update is responsible for incrementing the vote summary of the provided post
// objects according to the values provided with the given vote objects.
func Update(pos *poststorage.Object, vot *votestorage.Object) *poststorage.Object {
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
