package votereconciler

import (
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

	// DeleteVote prepares the provided vote objects so that they can be removed
	// from the underlying storage.
	//
	//     inp[0] the vote objects to be deleted
	//     out[0] the vote objects ready to be removed from storage
	//
	DeleteVote([]*votestorage.Object) ([]*votestorage.Object, error)

	// UpdateHash modifies the transaction hash of the votes as provided by the
	// given vote objects. Note that transaction hashes can only be updated once,
	// if their prior value was empty.
	//
	//     inp[0] the vote objects of the votes to modify
	//     inp[1] the transaction hashes to set as the new values
	//     out[0] the vote objects reflecting their updated state
	//
	UpdateHash([]*votestorage.Object, []string) ([]*votestorage.Object, error)
}
