package walletreconciler

import "github.com/uvio-network/apiserver/pkg/storage/walletstorage"

type Interface interface {
	// CreateWallet prepares the provided wallet objects so that they can be
	// persisted in the underlying storage the first time.
	//
	//     inp[0] the wallet objects according to the raw user input
	//     out[0] the wallet objects verified, extended and ready for storage
	//
	CreateWallet([]*walletstorage.Object) ([]*walletstorage.Object, error)
}
