package votestorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// CreateVote persists new vote objects in the underlying storage.
	//
	//     @inp[0] the vote objects to create
	//
	CreateVote([]*Object) error

	// SearchClaim returns the vote objects matching the given claim IDs.
	//
	//     @inp[0] the claim IDs used to search vote objects
	//     @out[0] the list of vote objects matching the given claim IDs
	//
	SearchClaim([]objectid.ID) ([]*Object, error)

	// SearchOwner returns the vote objects matching the given user IDs.
	//
	//     @inp[0] the user IDs used to search vote objects
	//     @out[0] the list of vote objects matching the given user IDs
	//
	SearchOwner([]objectid.ID) ([]*Object, error)

	// SearchOwnerClaim returns the vote objects matching the given user and claim
	// IDs. This search operation returns the votes that have been cast by a
	// specific user on a specific claim.
	//
	//     @inp[0] the user IDs used to search vote objects
	//     @inp[0] the claim IDs used to search vote objects
	//     @out[0] the list of vote objects matching the given user and claim IDs
	//
	SearchOwnerClaim([]objectid.ID, []objectid.ID) ([]*Object, error)

	// SearchVote returns the vote objects matching the given vote IDs.
	//
	//     @inp[0] the vote IDs used to search vote objects
	//     @out[0] the list of vote objects matching the given vote IDs
	//
	SearchVote([]objectid.ID) ([]*Object, error)
}
