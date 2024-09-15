package claimemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Interface interface {
	// Create emits an event that intends to trigger the creation of a claim with
	// the given lifecycle phase.
	//
	//     inp[0] the latest claim ID within a tree from which to create the next claim
	//     inp[1] the desired lifecycle phase of the claim to create
	//
	Create(objectid.ID, objectlabel.DesiredLifecycle) error
}
