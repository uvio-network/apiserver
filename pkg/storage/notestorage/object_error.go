package notestorage

import (
	"github.com/xh3b4sd/tracer"
)

var NoteKindEmptyError = &tracer.Error{
	Kind: "NoteKindEmptyError",
	Desc: "The request expects the vote kind not to be empty. The vote kind was found to be empty. Therefore the request failed.",
}

var NoteOwnerEmptyError = &tracer.Error{
	Kind: "NoteOwnerEmptyError",
	Desc: "The request expects the vote owner not to be empty. The vote owner was found to be empty. Therefore the request failed.",
}
