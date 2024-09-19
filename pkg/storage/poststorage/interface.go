package poststorage

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Interface interface {
	// CreatePost persists new post objects in the underlying storage.
	//
	//     @inp[0] the post objects to create
	//
	CreatePost([]*Object) error

	// DeleteExpiry removes the referenced IDs of the given post objects from the
	// lifecycle specific PostExpiry index.
	//
	//     @inp[0] the post objects to remove from the PostExpiry index
	//
	DeleteExpiry(inp []*Object) error

	// DeletePost removes existing post objects from the underlying storage.
	// Deleting posts is effectively the reversed operation of CreatePost.
	//
	//     @inp[0] the post objects to delete
	//
	DeletePost([]*Object) error

	// SearchComment returns the post objects that represent market specific
	// comments. All of the comments on a claim can be fetched using this method.
	//
	//     @inp[0] the claim IDs to search for
	//     @out[0] the list of post objects representing market specific comments
	//
	SearchComment([]objectid.ID) ([]*Object, error)

	// SearchCreated returns the post objects within the given pagination range in
	// reversed order of claim creation time. Given the created posts [A B C D E],
	// the first page [0 3] returns the last posts [C D E].
	//
	//     @inp[0] the start paging pointer defining the beginning of the page
	//     @inp[1] the stop paging pointer defining the end of the page
	//     @out[0] the list of post objects within the given pagination range
	//
	SearchCreated(int, int) ([]*Object, error)

	// SearchExpiry returns the lifecycle specific post objects recorded to have
	// expired already at the time of execution.
	//
	//     @out[0] the list of post objects already expired
	//
	SearchExpiry(objectlabel.DesiredLifecycle) ([]*Object, error)

	// SearchLabel returns the post objects grouped under all of the given
	// category labels. Multiple searches can be done for a set of labels each,
	// where each set of labels defines the intersection of claims to search for.
	//
	//     @inp[0] the list of category labels to search for
	//     @out[0] the list of post objects matching the given post IDs
	//
	SearchLabel([][]string) ([]*Object, error)

	// SearchLifecycle returns all the claims that have the specified lifecycle
	// phases defined.
	//
	//     @inp[0] the lifecycle phases to search for
	//     @out[0] the list of post objects having the specified lifecycle phases defined
	//
	SearchLifecycle([]string) ([]*Object, error)

	// SearchOwner returns the post objects created by the given user.
	//
	//     @inp[0] the user IDs to search for
	//     @out[0] the list of post objects created by the given user
	//
	SearchOwner([]objectid.ID) ([]*Object, error)

	// SearchOwnerComment returns the post objects that represent market specific
	// comments made by a specific user. If a user commented on a claim on which
	// they staked reputation or verified events, then SearchOwnerComment will
	// return all comments made by the given user on markets in which they
	// participated in.
	//
	//     @inp[0] the user IDs to search for
	//     @inp[0] the claim IDs to search for
	//     @out[0] the list of post objects representing market and user specific comments
	//
	SearchOwnerComment([]objectid.ID, []objectid.ID) ([]*Object, error)

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
