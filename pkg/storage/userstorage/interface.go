package userstorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// Create persists a new user object given the provided subject claim, if none
	// does exist already. Create is therefore idempotent and yields the same
	// persisted user object given the same provided subject claim.
	//
	//     @inp[0] the user object providing user specific information
	//     @out[0] the user object mapped to the given subject claim
	//
	Create(*Object) (*Object, error)

	// SearchSubject returns the user object mapped to the given subject claim, it
	// it exists. SearchSubject will return an error if there is no user mapping
	// already persisted between the external subject claim and the internal user
	// ID.
	//
	//     @inp[0] the external subject claim mapped to our internal user ID
	//     @out[0] the user object mapped to the given subject claim
	//
	SearchSubject(string) (*Object, error)

	// SearchUser returns the user objects matching the given user IDs.
	//
	//     @inp[0] the user IDs to search for
	//     @out[0] the list of user objects matching the given user IDs
	//
	SearchUser([]objectid.ID) ([]*Object, error)

	// UpdateUser modifies the given user objects in the underlying storage.
	//
	//     @inp[0] the user objects to update
	//
	UpdateUser([]*Object) error
}
