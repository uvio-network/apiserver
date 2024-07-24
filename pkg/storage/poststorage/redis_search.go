package poststorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchPost(use objectid.ID, inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(objectid.Fmt(inp, storageformat.PostObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(PostObjectNotFoundError, "%v", inp)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var out []*Object
	for i := range jsn {
		var obj *Object
		{
			obj = &Object{}
		}

		if jsn[i] != "" {
			err = json.Unmarshal([]byte(jsn[i]), obj)
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

func (r *Redis) SearchTree(use objectid.ID, inp []objectid.ID) ([]*Object, error) {
	var err error

	var out []*Object
	for _, x := range inp {
		// val will result in a list of all post IDs belonging to the given tree ID,
		// if any.
		var val []string
		{
			val, err = r.red.Sorted().Search().Order(posTre(x), 0, -1)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		// There might not be any post IDs, so we do not proceed, but instead
		// continue with the next tree ID, if any.
		if len(val) == 0 {
			continue
		}

		{
			lis, err := r.SearchPost(x, objectid.IDs(val))
			if err != nil {
				return nil, tracer.Mask(err)
			}

			out = append(out, lis...)
		}
	}

	return out, nil
}
