package walletstorage

import "github.com/xh3b4sd/objectid"

type Interface interface {
	// CreateWallet persists new wallet objects in the underlying storage.
	//
	//     @inp[0] the wallet objects to create
	//
	CreateWallet([]*Object) error

	// SearchAddress returns the wallet objects for the given addresses.
	//
	//     @inp[0] the wallet addresses to search for
	//     @out[0] the list of wallet objects for the given addresses
	//
	SearchAddress([]string) ([]*Object, error)

	// SearchOwner returns the wallet objects created by the given user.
	//
	//     @inp[0] the user IDs to search for
	//     @out[0] the list of wallet objects created by the given user
	//
	SearchOwner([]objectid.ID) ([]*Object, error)

	// SearchWallet returns the wallet objects matching the given wallet IDs.
	//
	//     @inp[0] the wallet IDs used to search wallet objects
	//     @out[0] the list of wallet objects matching the given wallet IDs
	//
	SearchWallet([]objectid.ID) ([]*Object, error)

	// UpdateWallet modifies the given wallet objects in the underlying storage.
	//
	//     @inp[0] the wallet objects to update
	//
	UpdateWallet([]*Object) error
}
