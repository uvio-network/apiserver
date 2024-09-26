package redis

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	var red redigo.Interface
	{
		red = redigo.Default()
	}

	var sto storage.Interface
	{
		sto = storage.New(storage.Config{
			Log: logger.Default(),
			Red: red,
		})
	}

	// --------------------------------------------------------------------- //

	//
	//     ./apiserver redis cla exp
	//

	var cla objectid.ID
	{
		cla = objectid.ID(arg[0])
	}

	var exp time.Time
	{
		exp = converter.StringToTime(arg[1])
	}

	// --------------------------------------------------------------------- //

	var pos *poststorage.Object
	{
		pos, err = r.seaPos(sto, cla)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = sto.Post().DeleteExpiry([]*poststorage.Object{pos})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		pos.Expiry = exp
	}

	{
		err = sto.Post().CreateExpiry([]*poststorage.Object{pos})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = sto.Post().UpdatePost([]*poststorage.Object{pos})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}

func (r *run) seaPos(sto storage.Interface, cla objectid.ID) (*poststorage.Object, error) {
	var err error

	var pos []*poststorage.Object
	{
		pos, err = sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(pos) != 1 {
		return nil, tracer.Mask(runtime.ExecutionFailedError)
	}

	return pos[0], nil
}
