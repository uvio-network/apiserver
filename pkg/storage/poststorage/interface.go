package poststorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// Create persists new post objects.
	//
	//     @inp[0] the post objects to create
	//     @out[0] the post objects mapped to their internal post IDs
	//
	Create([]*Object) ([]*Object, error)

	// SearchPost returns the post objects matching the given post IDs.
	//
	//     @inp[0] the calling user
	//     @inp[1] the post IDs to search for
	//     @out[0] the list of post objects matching the given post IDs
	//
	SearchPost(objectid.ID, []objectid.ID) ([]*Object, error)

	// SearchTree returns the post objects belonging to the given tree IDs.
	//
	//     @inp[0] the calling user
	//     @inp[1] the tree IDs to search for
	//     @out[0] the list of post objects belonging to the given tree IDs
	//
	SearchTree(objectid.ID, []objectid.ID) ([]*Object, error)
}