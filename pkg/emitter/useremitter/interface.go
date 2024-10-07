package useremitter

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	// UserCreate emits an event that signals the creation of a new user object.
	//
	//     inp[0] the user ID of the user that just registered
	//
	UserCreate(objectid.ID) error
}
