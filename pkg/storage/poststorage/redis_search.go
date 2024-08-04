package poststorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchLabels(inp [][]string) ([]*Object, error) {
	// cla will result in a list of all claim IDs grouped under all of the given
	// label names, if any.
	var cla []string
	for _, x := range inp {
		val, err := r.red.Sorted().Search().Inter(generic.Fmt(x, storageformat.PostLabel)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		cla = append(cla, val...)
	}

	// There might not be any claim IDs, so we do not proceed, but instead return
	// nothing.
	if len(cla) == 0 {
		return nil, nil
	}

	// Having collected all claim IDs, we go ahead and search all claim objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchPost(objectid.IDs(cla))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redis) SearchPage(beg int, end int) ([]*Object, error) {
	var err error

	// val will result in a list of all post IDs within the given pagination
	// range, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Order(storageformat.PostCreated, -(end + 1), -(beg + 1))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any post IDs, so we do not proceed, but instead return
	// nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		lis, err := r.SearchPost(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redis) SearchPost(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Fmt(inp, storageformat.PostObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(PostObjectNotFoundError, "%v", inp)
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

func (r *Redis) SearchTree(inp []objectid.ID) ([]*Object, error) {
	// pos will result in a list of all post IDs belonging to the given tree ID,
	// if any.
	var pos []string
	for _, x := range inp {
		val, err := r.red.Sorted().Search().Order(posTre(x), 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		pos = append(pos, val...)
	}

	// There might not be any post IDs, so we do not proceed, but instead return
	// nothing.
	if len(pos) == 0 {
		return nil, nil
	}

	// Having collected all post IDs, we go ahead and search all post objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchPost(objectid.IDs(pos))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}
