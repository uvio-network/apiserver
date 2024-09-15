package uvxemitter

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// Mint emits an event that intends to trigger the minting of UVX tokens for
	// the given user ID.
	//
	//     inp[0] the user ID to mint tokens for
	//
	Mint(objectid.ID) error
}
