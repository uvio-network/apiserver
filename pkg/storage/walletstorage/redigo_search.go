package walletstorage

import (
	"encoding/json"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/redigo/simple"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) SearchAddress(add []string) ([]*Object, error) {
	var err error

	// wal will result in a list of all wallet IDs associated to the given wallet
	// addresses, if any.
	var wal []string
	{
		wal, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.WalletAddress, add)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any wallet IDs, so we do not proceed, but instead return
	// nothing.
	if len(wal) == 0 {
		return nil, nil
	}

	// Having collected all wallet IDs, we go ahead and search all wallet objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchWallet(objectid.IDs(wal))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redigo) SearchOwner(own []objectid.ID) ([]*Object, error) {
	var err error

	// wal will result in a list of all wallet IDs belonging to the given user
	// IDs, if any.
	var wal []string
	{
		wal, err = r.red.Sorted().Search().Union(generic.Arg1(storageformat.WalletOwner, own)...)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// There might not be any wallet IDs, so we do not proceed, but instead return
	// nothing.
	if len(wal) == 0 {
		return nil, nil
	}

	// Having collected all wallet IDs, we go ahead and search all wallet objects at
	// once.
	var out []*Object
	{
		lis, err := r.SearchWallet(objectid.IDs(wal))
		if err != nil {
			return nil, tracer.Mask(err)
		}

		out = append(out, lis...)
	}

	return out, nil
}

func (r *Redigo) SearchWallet(inp []objectid.ID) ([]*Object, error) {
	var err error

	var jsn []string
	{
		jsn, err = r.red.Simple().Search().Multi(generic.Arg1(storageformat.WalletObject, inp)...)
		if simple.IsNotFound(err) {
			return nil, tracer.Maskf(WalletObjectNotFoundError, "%v", inp)
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
