package claimemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

type Interface interface {
	// BalanceUpdate emits an event that intends to trigger the update of a
	// claim's user balances.
	//
	//     inp[0] the block number at which the call to BalanceUpdate was made
	//     inp[1] the claim ID of the expired resolve
	//
	BalanceUpdate(uint64, objectid.ID) error

	// ResolveCreate emits an event that intends to trigger the creation of a
	// resolve using the given claim ID, which is from either a propose or
	// dispute.
	//
	//     inp[0] the block number at which the call to ResolveCreate was made
	//     inp[1] the claim ID of either the expired propose or dispute
	//
	ResolveCreate(uint64, objectid.ID) error
}
