package notestorage

import (
	"github.com/xh3b4sd/tracer"
)

var NoteKindEmptyError = &tracer.Error{
	Kind: "NoteKindEmptyError",
	Desc: "The request expects the note kind not to be empty. The note kind was found to be empty. Therefore the request failed.",
}

var NoteOwnerEmptyError = &tracer.Error{
	Kind: "NoteOwnerEmptyError",
	Desc: "The request expects the note owner not to be empty. The note owner was found to be empty. Therefore the request failed.",
}
