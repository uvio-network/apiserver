package votestorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) SearchClaim(cla []objectid.ID) ([]*Object, error) {
	// vot will result in a list of all vote IDs belonging to the given claim IDs,
	// if any.
	var vot []string
	for _, x := range cla {
		val, err := r.red.Sorted().Search().Order(votCla(x), 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		vot = append(vot, val...)
	}

	// There might not be any vote IDs, so we do not proceed, but instead return
	// nothing.
	if len(vot) == 0 {
		return nil, nil
	}

	// Having collected all vote IDs, we go ahead and search all vote objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchVote(objectid.IDs(vot))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redis) SearchOwner(own []objectid.ID) ([]*Object, error) {
	// vot will result in a list of all vote IDs belonging to the given user IDs,
	// if any.
	var vot []string
	for _, x := range own {
		val, err := r.red.Sorted().Search().Order(votOwn(x), 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		vot = append(vot, val...)
	}

	// There might not be any vote IDs, so we do not proceed, but instead return
	// nothing.
	if len(vot) == 0 {
		return nil, nil
	}

	// Having collected all vote IDs, we go ahead and search all vote objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchVote(objectid.IDs(vot))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redis) SearchOwnerClaim(own []objectid.ID, cla []objectid.ID) ([]*Object, error) {
	if len(own) != len(cla) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(own), len(cla))
	}

	// vot will result in a list of all vote IDs belonging to the given user and
	// claim IDs, if any.
	var vot []string
	for i := range own {
		val, err := r.red.Sorted().Search().Order(votOwnCla(own[i], cla[i]), 0, -1)
		if err != nil {
			return nil, tracer.Mask(err)
		}

		vot = append(vot, val...)
	}

	// There might not be any vote IDs, so we do not proceed, but instead return
	// nothing.
	if len(vot) == 0 {
		return nil, nil
	}

	// Having collected all vote IDs, we go ahead and search all vote objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchVote(objectid.IDs(vot))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redis) SearchVote(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Fmt(inp, storageformat.VoteObject)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(VoteObjectNotFoundError, "%v", inp)
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
