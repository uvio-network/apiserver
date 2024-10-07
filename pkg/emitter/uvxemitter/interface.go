package uvxemitter

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// UvxMint emits an event that intends to trigger the minting of UVX tokens
	// for the given user ID.
	//
	//     inp[0] the user ID to mint UVX tokens for
	//
	UvxMint(objectid.ID) error
}
