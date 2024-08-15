package votereconciler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
)

type Interface interface {
	// CreateVote prepares the provided vote objects so that they can be persisted
	// in the underlying storage the first time.
	//
	//     inp[0] the vote objects according to the raw user input
	//     out[0] the vote objects verified, extended and ready for storage
	//
	CreateVote([]*votestorage.Object) ([]*votestorage.Object, error)

	// UpdateHash modifies the transaction hash of the votes as provided by the
	// given object IDs. Note that transaction hashes can only be updated once, if
	// their prior value was empty.
	//
	//     inp[0] the calling user ID for ownership verification
	//     inp[1] the object IDs of the votes to modify
	//     inp[2] the transaction hashes to set as the new values
	//     out[0] the vote objects reflecting their updated transaction hashes
	//
	UpdateHash(objectid.ID, []objectid.ID, []string) ([]*votestorage.Object, error)
}
