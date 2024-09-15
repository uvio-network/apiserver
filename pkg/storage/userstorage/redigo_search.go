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
