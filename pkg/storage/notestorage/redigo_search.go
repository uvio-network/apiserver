package notestorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) SearchNote(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Arg1(storageformat.NoteObject, inp)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(NoteObjectNotFoundError, "%v", inp)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	for _, x := range jsn {
		var obj *Object
		{
			obj = &Object{}
		}

		if x != "" {
			err = json.Unmarshal([]byte(x), obj)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		{
			out = append(out, obj)
		}
	}

	return out, nil
}

func (r *Redigo) SearchPage(use objectid.ID, kin []string, beg int, end int) ([]*Object, error) {
	// val will result in a list of all note IDs within the given pagination
	// range, if any.
	var val []string
	for _, x := range kin {
		lis, err := r.red.Sorted().Search().Order(notOwnKin(use, x), -(end + 1), -(beg + 1))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = append(val, lis...)
	}

	// There might not be any note IDs, so we do not proceed, but instead return
	// nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		lis, err := r.SearchNote(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redigo) SearchTime(use objectid.ID, kin []string, beg int64, end int64) ([]*Object, error) {
	// val will result in a list of all note IDs within the given creation
	// timestamps, if any.
	var val []string
	for _, x := range kin {
		lis, err := r.red.Sorted().Search().Score(notOwnKin(use, x), float64(beg), float64(end))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		val = append(val, lis...)
	}

	// There might not be any note IDs, so we do not proceed, but instead return
	// nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		lis, err := r.SearchNote(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}
