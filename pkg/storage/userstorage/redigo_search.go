package userstorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) SearchReputation(beg int, end int) ([]*Object, error) {
	var err error

	// val will result in a list of all user IDs within the given pagination
	// range, if any.
	var val []string
	{
		val, err = r.red.Sorted().Search().Order(storageformat.UserReputation, -(end + 1), -(beg + 1))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any user IDs, so we do not proceed, but instead return
	// nothing.
	if len(val) == 0 {
		return nil, nil
	}

	var out []*Object
	{
		lis, err := r.SearchUser(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redigo) SearchSubject(sub string) (*Object, error) {
	var err error

	if sub == "" {
		return nil, tracer.Mask(UserSubjectEmptyError)
	}

	var val []string
	{
		val, err = r.red.Simple().Search().Multi(useSub(sub))
		if simple.IsNotFound(err) {
			return nil, tracer.Mask(SubjectClaimMappingError)
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(val) != 1 {
		return nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	var out []*Object
	{
		out, err = r.SearchUser(objectid.IDs(val))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return out[0], nil
}

func (r *Redigo) SearchUser(use []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Arg1(storageformat.UserObject, use)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(UserNotFoundError, "%v", use)
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
