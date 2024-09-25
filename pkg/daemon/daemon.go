package daemon

import (
	"crypto/rand"
	"net"
	"time"

	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/breakr"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/locker/lock"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/redigo/pool"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/rescue/engine"
	"github.com/xh3b4sd/tracer"
)

type Daemon struct {
	emi emitter.Interface
	env envvar.Env
	lis net.Listener
	loc locker.Interface
	log logger.Interface
	rec reconciler.Interface
	red redigo.Interface
	res rescue.Interface
	sam *sample.Sample
	sto storage.Interface
}

func New(env envvar.Env) *Daemon {
	var err error

	var log logger.Interface
	{
		log = logger.New(logger.Config{
			Filter: logger.NewLevelFilter(env.LogLevel),
		})
	}

	var lis net.Listener
	{
		lis, err = net.Listen("tcp", net.JoinHostPort(env.HttpHost, env.HttpPort))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var red redigo.Interface
	{
		red = redigo.Default()
	}

	var res rescue.Interface
	{
		res = engine.New(engine.Config{
			Logger: log,
			Queue:  "api.uvio.network", // rescue.io/api.uvio.network
			Redigo: red,
			Sepkey: "/",
		})
	}

	var emi emitter.Interface
	{
		emi = emitter.New(emitter.Config{
			Log: log,
			Res: res,
		})
	}

	var loc locker.Interface
	{
		loc = defLoc(red.Listen())
	}

	var sto storage.Interface
	{
		sto = storage.New(storage.Config{
			Log: log,
			Red: red,
		})
	}

	var rec reconciler.Interface
	{
		rec = reconciler.New(reconciler.Config{
			Log: log,
			Red: red,
			Sto: sto,
		})
	}

	var sam *sample.Sample
	{
		sam = sample.New(sample.Config{
			Rea: rand.Reader,
		})
	}

	return &Daemon{
		emi: emi,
		env: env,
		lis: lis,
		loc: loc,
		log: log,
		rec: rec,
		red: red,
		res: res,
		sam: sam,
		sto: sto,
	}
}

func (d *Daemon) Env() envvar.Env {
	return d.env
}

func (d *Daemon) Rec() reconciler.Interface {
	return d.rec
}

func (d *Daemon) Sto() storage.Interface {
	return d.sto
}

func defLoc(add string) locker.Interface {
	return lock.New(lock.Config{
		Brk: breakr.New(breakr.Config{
			Failure: breakr.Failure{
				Budget: 30,
				Cooler: 1 * time.Second,
			},
		}),
		Poo: pool.NewSinglePoolWithAddress(add),
	})
}
