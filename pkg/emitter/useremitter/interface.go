package useremitter

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// CompetenceUpdate emits an event that intends to trigger the update of the
	// competence specific user metrics for all users that participated in the
	// given market.
	//
	//     inp[0] the claim ID of the settled balance
	//
	CompetenceUpdate(objectid.ID) error

	// IntegritiyUpdate emits an event that intends to trigger the update of the
	// integritiy specific user metrics for all users that participated in the
	// given market.
	//
	//     inp[0] the claim ID of the settled balance
	//
	IntegrityUpdate(objectid.ID) error

	// ReputationSearch emits an event that signals the lookup for privileged
	// voter addresses.
	//
	//     inp[0] the post ID of the dispute that has been expired
	//
	ReputationSearch(objectid.ID) error

	// UserCreate emits an event that signals the creation of a new user object.
	//
	//     inp[0] the user ID of the user that just registered
	//
	UserCreate(objectid.ID) error
}
