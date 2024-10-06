package replicate

import (
	"net"

	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	var env envvar.Env
	{
		env = envvar.Load("sepolia")
	}

	// This is the remote Redis we are reading from.
	var rea redigo.Interface
	{
		c := redigo.Config{
			Kind:    redigo.KindSingle,
			Address: net.JoinHostPort(env.RedisHost, "6379"),
			User:    env.RedisUser,
			Pass:    env.RedisPass,
		}

		rea, err = redigo.New(c)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// This is the local Redis we are writing to.
	var wri redigo.Interface
	{
		wri = redigo.Default()
	}

	// --------------------------------------------------------------------- //

	dnc := make(chan struct{}, 1)
	erc := make(chan error, 1)
	stc := make(chan string, 1)

	go func() {
		defer close(dnc)
		defer close(erc)
		defer close(stc)

		err = rea.Walker().Search().Keys("*", dnc, stc)
		if err != nil {
			erc <- tracer.Mask(err)
			return
		}
	}()

	go func() {
		for x := range stc {
			var typ string
			{
				typ, err = rea.Walker().Search().Type(x)
				if err != nil {
					erc <- tracer.Mask(err)
					return
				}
			}

			if typ == "string" {
				var val []string
				{
					val, err = rea.Simple().Search().Multi(x)
					if err != nil {
						erc <- tracer.Mask(err)
						return
					}
				}

				if len(val) != 1 {
					erc <- tracer.Mask(runtime.ExecutionFailedError)
					return
				}

				{
					err = wri.Simple().Create().Element(x, val[0])
					if err != nil {
						erc <- tracer.Mask(err)
						return
					}
				}
			}

			if typ == "zset" {
				var val []string
				{
					val, err = rea.Sorted().Search().Order(x, 0, -1, true)
					if err != nil {
						erc <- tracer.Mask(err)
						return
					}
				}

				if len(val) == 0 {
					erc <- tracer.Mask(runtime.ExecutionFailedError)
					return
				}

				for i := 0; i < len(val); i += 2 {
					err = wri.Sorted().Create().Score(x, val[i], converter.StringToFloat(val[i+1]))
					if err != nil {
						erc <- tracer.Mask(err)
						return
					}
				}
			}
		}
	}()

	{
		err = <-erc
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
