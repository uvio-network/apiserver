package notecreatehandler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *InternHandler) Ensure(tas *task.Task) error {
	var err error

	var kin string
	{
		kin = tas.Meta.Get(objectlabel.NoteKind)
	}

	if kin == "resolveCreate" {
		err = h.resolveCreate(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) resolveCreate(tas *task.Task) error {
	var err error

	var res *poststorage.Object
	{
		res, err = h.searchClaim(tas)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var not []*notestorage.Object
	for _, v := range res.Samples {
		not = append(not, &notestorage.Object{
			Kind:     objectlabel.NoteKindVoterSelected,
			Owner:    objectid.ID(v),
			Resource: res.ID,
		})
	}

	{
		err = h.sto.Note().CreateNote(not)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (h *InternHandler) searchClaim(tas *task.Task) (*poststorage.Object, error) {
	var err error

	var cla objectid.ID
	{
		cla = objectid.ID(tas.Meta.Get(objectlabel.NoteResource))
	}

	var pos poststorage.Slicer
	{
		pos, err = h.sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != 1 {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "expected exactly one post object for ID %s", cla)
	}

	return pos[0], nil
}
