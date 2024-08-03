package votevalidator

import "github.com/uvio-network/apiserver/pkg/storage/votestorage"

type Interface interface {
	// CreateVote prepares the provided vote objects so that they can be persisted
	// in the underlying storage the first time.
	//
	//     inp[0] the vote objects according to the raw user input
	//     out[0] the vote objects verified, extended and ready for storage
	//
	CreateVote([]*votestorage.Object) ([]*votestorage.Object, error)
}
