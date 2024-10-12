package claimemitter

import (
	"github.com/xh3b4sd/objectid"
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
	// dispute. For the final dispute the event can be created with the optional
	// meta data of privileged staker addresses, which should be selected based on
	// their reputation metrics.
	//
	//     inp[0] the block number at which the call to ResolveCreate was made
	//     inp[1] the claim ID of either the expired propose or dispute
	//     inp[2] the list of staker indices to select as voters, if any
	//
	ResolveCreate(uint64, objectid.ID, ...string) error
}
