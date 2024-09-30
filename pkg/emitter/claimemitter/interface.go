package claimemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Interface interface {
	// Create emits an event that intends to trigger the creation of a claim with
	// the given lifecycle phase.
	//
	//     inp[0] the block number at which the call to Create was made
	//     inp[1] the latest claim ID within a tree from which to create the next claim
	//     inp[2] the desired lifecycle phase of the claim to create
	//
	Create(uint64, objectid.ID, objectlabel.DesiredLifecycle) error

	// Update emits an event that intends to trigger the update of a claim's
	// lifecycle phase.
	//
	//     inp[0] the block number at which the call to Update was made
	//     inp[1] the propose ID for which to update the lifecycle phase
	//     inp[2] the lifecycle phase to update
	//
	Update(uint64, objectid.ID, objectlabel.DesiredLifecycle) error
}
