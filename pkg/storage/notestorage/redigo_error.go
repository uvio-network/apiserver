package notestorage

import (
	"github.com/xh3b4sd/tracer"
)

var NoteObjectNotFoundError = &tracer.Error{
	Kind: "NoteObjectNotFoundError",
	Desc: "The request expects a note object to exist. The note object was not found to exist. Therefore the request failed.",
}
