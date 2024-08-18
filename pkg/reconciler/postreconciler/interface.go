package postreconciler

import (
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
)

type Interface interface {
	// CreatePost prepares the provided post objects so that they can be persisted
	// in the underlying storage the first time.
	//
	//     inp[0] the post objects according to the raw user input
	//     out[0] the post objects verified, extended and ready for storage
	//
	CreatePost([]*poststorage.Object) ([]*poststorage.Object, error)

	// UpdateHash modifies the transaction hash of the claims as provided by the
	// given post objects. Note that transaction hashes can only be updated once,
	// if their prior value was empty.
	//
	//     inp[0] the post objects of the claims to modify
	//     inp[1] the transaction hashes to set as the new values
	//     out[0] the post objects reflecting their updated state
	//
	UpdateHash([]*poststorage.Object, []string) ([]*poststorage.Object, error)

	// UpdateMeta modifies the post meta of the provided post objects.
	//
	//     inp[0] the post objects of the claims to modify
	//     inp[1] the meta data to set as the new values
	//     out[0] the post objects reflecting their updated state
	//
	UpdateMeta([]*poststorage.Object, []string) ([]*poststorage.Object, error)

	// UpdateVotes modifies the vote summary of the existing post objects that are
	// linked to the provided vote objects, so that said vote summary properly
	// reflects the votes being cast.
	//
	//     inp[0] the vote objects of the votes being cast
	//     out[0] the post objects reflecting their updated vote summary
	//
	UpdateVotes([]*votestorage.Object) ([]*poststorage.Object, error)
}
