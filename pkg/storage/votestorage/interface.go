package votestorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// CreateVote persists new vote objects in the underlying storage.
	//
	//     @inp[0] the vote objects to create
	//
	CreateVote([]*Object) error

	// SearchVote returns the vote objects matching the given vote IDs.
	//
	//     @inp[0] the vote IDs to search for
	//     @out[0] the list of vote objects matching the given vote IDs
	//
	SearchVote([]objectid.ID) ([]*Object, error)
}
