package poststorage

import (
	"encoding/json"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) SearchComment(cla []objectid.ID) ([]*Object, error) {
	var err error

	// com will result in a list of all comment IDs belonging to the given claim
	// IDs, if any.
	var com []string
	{
		com, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.PostComment, cla)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any comment IDs, so we do not proceed, but instead
	// return nothing.
	if len(com) == 0 {
		return nil, nil
	}

	// Having collected all comment IDs, we go ahead and search all comment
	// objects at once.
	var out []*Object
	{
		lis, err := r.SearchPost(objectid.IDs(com))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redigo) SearchExpiry(lif objectlabel.DesiredLifecycle) ([]*Object, error) {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	// val will result in a list of all post IDs within the given pagination
	// range, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Score(posExp(lif), 0, float64(now.Unix()))
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

func (r *Redigo) SearchLabel(inp [][]string) ([]*Object, error) {
	// cla will result in a list of all claim IDs grouped under all of the given
	// label names, if any.
	var cla []string
	for _, x := range inp {
		val, err := r.red.Sorted().Search().Inter(generic.Arg1(storageformat.PostLabel, x)...)
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

func (r *Redigo) SearchLifecycle(lif []objectlabel.DesiredLifecycle) ([]*Object, error) {
	var err error

	var key []string
	{
		key = generic.Func(lif, func(x objectlabel.DesiredLifecycle) string {
			// When we show settled claims, then we show the latest lifecycle phase
			// there is. This final lifecycle phase has no expiry. That is why we use
			// the storage key of the static sorted set that contains all claim IDs
			// per lifecycle, regardless of the claim's expiry.
			if x == objectlabel.LifecycleBalance {
				return posLif(x)
			}

			// When we show ongoing claims, then we only want to show those claims
			// that are current in the requested lifecycle phase. For that reason we
			// use the storage key of expiring claims per lifecycle phase.
			return posExp(x)
		})
	}

	// cla will result in a list of all claim IDs defining the given claim
	// lifecycle, if any.
	var cla []string
	{
		cla, err = r.red.Sorted().Search().Union(key...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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

func (r *Redigo) SearchOwner(use []objectid.ID) ([]*Object, error) {
	var err error

	// pos will result in a list of all post IDs created by the given user IDs, if
	// any.
	var pos []string
	{
		pos, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.PostOwner, use)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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

func (r *Redigo) SearchOwnerComment(own []objectid.ID, cla []objectid.ID) ([]*Object, error) {
	var err error

	if len(own) != len(cla) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(own), len(cla))
	}

	// com will result in a list of all comment IDs belonging to the given user
	// and claim IDs, if any.
	var com []string
	{
		com, err = r.red.Sorted().Search().Union(generic.Arg2(storageformat.PostOwnerComment, own, cla)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any comment IDs, so we do not proceed, but instead
	// return nothing.
	if len(com) == 0 {
		return nil, nil
	}

	// Having collected all comment IDs, we go ahead and search all comment
	// objects at once.
	var out []*Object
	{
		lis, err := r.SearchPost(objectid.IDs(com))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redigo) SearchPage(beg int, end int) ([]*Object, error) {
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

func (r *Redigo) SearchPost(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Arg1(storageformat.PostObject, inp)...)
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

func (r *Redigo) SearchTree(tre []objectid.ID) ([]*Object, error) {
	var err error

	// pos will result in a list of all post IDs belonging to the given tree IDs,
	// if any.
	var pos []string
	{
		pos, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.PostTree, tre)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
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
