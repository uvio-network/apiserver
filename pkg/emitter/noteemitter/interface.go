package noteemitter

import "github.com/xh3b4sd/objectid"

type Interface interface {
	// NoteCreate emits an event that intends to create notifications for specific
	// users.
	//
	//     inp[0] the note kind to create
	//     inp[1] the note message to create
	//     inp[2] the note resource to create
	//
	NoteCreate(kin string, mes string, res objectid.ID) error
}
