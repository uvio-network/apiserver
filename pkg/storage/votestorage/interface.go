package votestorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// Create persists new vote objects.
	//
	//     @inp[0] the vote objects to create
	//     @out[0] the vote objects mapped to their internal vote IDs
	//
	Create([]*Object) ([]*Object, error)

	// SearchVote returns the vote objects matching the given vote IDs.
	//
	//     @inp[0] the calling user
	//     @inp[1] the vote IDs to search for
	//     @out[0] the list of vote objects matching the given vote IDs
	//
	SearchVote(objectid.ID, []objectid.ID) ([]*Object, error)
}
