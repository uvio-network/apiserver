package poststorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// CreatePost persists new post objects in the underlying storage.
	//
	//     @inp[0] the post objects to create
	//
	CreatePost([]*Object) error

	// SearchComment returns the post objects that represent market specific
	// comments. If a user commented on a claim on which they staked reputation or
	// verified events, then SearchComment will return all comments made by the
	// given user on markets in which they participated in.
	//
	//     @inp[0] the user IDs to search for
	//     @inp[0] the claim IDs to search for
	//     @out[0] the list of post objects representing market specific comments
	//
	SearchComment([]objectid.ID, []objectid.ID) ([]*Object, error)

	// SearchLabel returns the post objects grouped under all of the given
	// category labels. Multiple searches can be done for a set of labels each,
	// where each set of labels defines the intersection of claims to search for.
	//
	//     @inp[0] the list of category labels to search for
	//     @out[0] the list of post objects matching the given post IDs
	//
	SearchLabel([][]string) ([]*Object, error)

	// SearchPage returns the post objects within the given pagination range in
	// reversed order of creation time. Given the chronologically persisted posts
	// [A B C D E], the first page [0 3] returns the last posts [C D E].
	//
	//     @inp[0] the start paging pointer defining the beginning of the page
	//     @inp[1] the stop paging pointer defining the end of the page
	//     @out[0] the list of post objects within the given pagination range
	//
	SearchPage(int, int) ([]*Object, error)

	// SearchPost returns the post objects matching the given post IDs.
	//
	//     @inp[0] the post IDs to search for
	//     @out[0] the list of post objects matching the given post IDs
	//
	SearchPost([]objectid.ID) ([]*Object, error)

	// SearchTree returns the post objects belonging to the given tree IDs.
	//
	//     @inp[0] the tree IDs to search for
	//     @out[0] the list of post objects belonging to the given tree IDs
	//
	SearchTree([]objectid.ID) ([]*Object, error)

	// UpdatePost modifies the given post objects in the underlying storage.
	//
	//     @inp[0] the post objects to update
	//
	UpdatePost([]*Object) error
}
