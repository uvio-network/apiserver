package poststorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// Create persists new post objects.
	//
	//     @inp[0] the post objects to create
	//     @out[0] the post objects mapped to their internal post IDs
	//
	Create([]*Object) ([]*Object, error)

	// SearchLabels returns the post objects grouped under all of the given
	// category labels. Multiple searches can be done for a set of labels each,
	// where each set of labels defines the intersection of claims to search for.
	//
	//     @inp[0] the calling user
	//     @inp[1] the list of category labels to search for
	//     @out[0] the list of post objects matching the given post IDs
	//
	SearchLabels(objectid.ID, [][]string) ([]*Object, error)

	// SearchPage returns the post objects within the given pagination range in
	// reversed order of creation time. Given the chronologically persisted posts
	// [A B C D E], the first page [0 3] returns the last posts [C D E].
	//
	//     @inp[0] the calling user
	//     @inp[1] the start paging pointer defining the beginning of the page
	//     @inp[2] the stop paging pointer defining the end of the page
	//     @out[0] the list of post objects within the given pagination range
	//
	SearchPage(objectid.ID, int, int) ([]*Object, error)

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
