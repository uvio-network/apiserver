package userstorage

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// UpdateCompetence adds the referenced IDs of the given user objects to the
	// reputation specific UserCompetence index. This form of indexing enables us
	// to search for the users with the most competence on the Uvio platform.
	//
	//     @inp[0] the user objects to add to the UserCompetence index
	//
	UpdateCompetence(inp []*Object) error

	// UpdateIntegrity adds the referenced IDs of the given user objects to the
	// reputation specific UserIntegrity index. This form of indexing enables us
	// to search for the users with the most integrity on the Uvio platform.
	//
	//     @inp[0] the user objects to add to the UserIntegrity index
	//
	UpdateIntegrity(inp []*Object) error

	// CreateUser persists a new user object given the provided subject claim, if
	// none does exist already. Create is therefore idempotent and yields the same
	// persisted user object given the same provided subject claim.
	//
	//     @inp[0] the user object providing user specific information
	//     @out[0] the user object mapped to the given subject claim
	//
	CreateUser(*Object) (*Object, error)

	// DeleteCompetence removes the lowest integrity users from the limited
	// integrity index.
	DeleteCompetence() error

	// DeleteIntegrity removes the lowest integrity users from the limited
	// integrity index.
	DeleteIntegrity() error

	// SearchReputation returns the user objects within the given pagination range in
	// reversed order of user reputation scores. Given the indexed users [A B C D E],
	// the first page [0 3] returns the highest reputation users [C D E].
	//
	//     @inp[0] the start paging pointer defining the beginning of the page
	//     @inp[1] the stop paging pointer defining the end of the page
	//     @out[0] the list of user objects within the given pagination range
	//
	SearchReputation(int, int) ([]*Object, error)

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
