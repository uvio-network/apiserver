package notestorage

import (
	"time"

	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) CreateNote(inp []*Object) error {
	var err error

	for i := range inp {
		{
			err := inp[i].Verify()
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var now time.Time
		{
			now = time.Now().UTC()
		}

		{
			inp[i].Created = now
			inp[i].ID = objectid.Random(objectid.Time(now))
		}

		// Create the normalized key-value pair for the note object. With that we
		// can search for note objects using their IDs.
		{
			err = r.red.Simple().Create().Element(notObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Persist note object IDs for every user and every topic respectively.
		// This mapping enables us to search for all notes that belong to any given
		// user on any given topic.
		{
			err = r.red.Sorted().Create().Score(notOwnKin(inp[i].Owner, inp[i].Kind), inp[i].ID.String(), float64(inp[i].Created.Unix()))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Cleanup the extraneous note objects in order to maintain a fixed size
		// notification log per user per topic. Above we add one note, below we
		// remove one in exchange above a certain threshold.
		{
			err = r.deleteNote(inp[i])
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
