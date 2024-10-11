package userstorage

import (
	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) DeleteReputation() error {
	var err error

	{
		err = r.red.Sorted().Delete().Limit(storageformat.UserReputation, 1000)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
