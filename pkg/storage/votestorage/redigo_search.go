package votestorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) SearchClaim(cla []objectid.ID) ([]*Object, error) {
	var err error

	// vot will result in a list of all vote IDs belonging to the given claim IDs,
	// if any.
	var vot []string
	{
		vot, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.VoteClaim, cla)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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

func (r *Redigo) SearchOwner(own []objectid.ID) ([]*Object, error) {
	var err error

	// vot will result in a list of all vote IDs belonging to the given user IDs,
	// if any.
	var vot []string
	{
		vot, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.VoteOwner, own)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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

func (r *Redigo) SearchOwnerClaim(own []objectid.ID, cla []objectid.ID) ([]*Object, error) {
	var err error

	if len(own) != len(cla) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(own), len(cla))
	}

	// vot will result in a list of all vote IDs belonging to the given user and
	// claim IDs, if any.
	var vot []string
	{
		vot, err = r.red.Sorted().Search().Union(generic.Arg2(storageformat.VoteOwnerClaim, own, cla)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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

func (r *Redigo) SearchVote(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Arg1(storageformat.VoteObject, inp)...)
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
