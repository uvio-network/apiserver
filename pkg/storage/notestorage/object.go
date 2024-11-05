package notestorage

import (
	"time"

	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Created time.Time   `json:"created,omitempty"`
	ID      objectid.ID `json:"id,omitempty"`
	Kind    string      `json:"kind,omitempty"`
	Message string      `json:"message,omitempty"`
	Owner   objectid.ID `json:"owner,omitempty"`
}

func (o *Object) Verify() error {
	{
		if o.Kind == "" {
			return tracer.Maskf(NoteKindEmptyError, o.Kind)
		}
	}

	{
		if o.Owner == "" {
			return tracer.Mask(NoteOwnerEmptyError)
		}
	}

	return nil
}
