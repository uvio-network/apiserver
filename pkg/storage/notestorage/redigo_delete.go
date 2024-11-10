package notestorage

import (
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) deleteNote(inp *Object) error {
	var err error

	for {
		var cou int64
		{
			cou, err = r.red.Sorted().Metric().Count(notOwn(inp.Owner))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if cou > 100 {
			// Every notification index is sorted by creation time in ascending order.
			// Below we get the oldest note object ID, which is the first item in the
			// queue.
			var val []string
			{
				val, err = r.red.Sorted().Search().Order(notOwn(inp.Owner), 0, 1)
				if err != nil {
					return tracer.Mask(err)
				}
			}

			if len(val) != 1 {
				return tracer.Maskf(runtime.ExecutionFailedError, "%d != 1", len(val))
			}

			var oid objectid.ID
			{
				oid = objectid.ID(val[0])
			}

			// Remove the oldest note object from the underlying storage.
			{
				_, err = r.red.Simple().Delete().Multi(notObj(oid))
				if err != nil {
					return tracer.Mask(err)
				}
			}

			// Remove the oldest note object ID from the very notification index that
			// we got it from in the first place.
			{
				err = r.red.Sorted().Delete().Score(notOwn(inp.Owner), oid.Float())
				if err != nil {
					return tracer.Mask(err)
				}
			}
		} else {
			// Once we guaranteed 100 notes in the given notification index, we stop
			// iterating. Under normal circumstances we should usually only iterate
			// once, and if we failed previously, we should only have to iterate twice
			// in order to get the job done.
			break
		}
	}

	return nil
}
