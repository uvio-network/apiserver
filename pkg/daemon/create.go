package daemon

import (
	"net"
	"time"

	"github.com/gorilla/mux"
	"github.com/twitchtv/twirp"
	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/server"
	"github.com/uvio-network/apiserver/pkg/server/interceptor/failedinterceptor"
	"github.com/uvio-network/apiserver/pkg/server/middleware/authmiddleware"
	"github.com/uvio-network/apiserver/pkg/server/middleware/corsmiddleware"
	"github.com/uvio-network/apiserver/pkg/server/middleware/usermiddleware"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/worker"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler"
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
	sto storage.Interface
}

func New(env envvar.Env) *Daemon {
	var err error

	var log logger.Interface
	{
		log = logger.Default()
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
			Sto: sto,
		})
	}

	// --------------------------------------------------------------------- //

	var dae *Daemon
	{
		dae = &Daemon{
			emi: emi,
			env: env,
			lis: lis,
			loc: loc,
			log: log,
			rec: rec,
			red: red,
			res: res,
			sto: sto,
		}
	}

	return dae
}

func (d *Daemon) Server() *server.Server {
	var shn *serverhandler.Handler
	{
		shn = serverhandler.New(serverhandler.Config{
			Emi: d.emi,
			Loc: d.loc,
			Log: d.log,
			Rec: d.rec,
			Sto: d.sto,
		})
	}

	var srv *server.Server
	{
		srv = server.New(server.Config{
			Han: shn.Hand(),
			Int: []twirp.Interceptor{
				failedinterceptor.NewInterceptor(failedinterceptor.InterceptorConfig{Log: d.log}).Interceptor,
			},
			Lis: d.lis,
			Log: d.log,
			Mid: []mux.MiddlewareFunc{
				corsmiddleware.NewMiddleware(corsmiddleware.MiddlewareConfig{Log: d.log}).Handler,
				authmiddleware.NewMiddleware(authmiddleware.MiddlewareConfig{Aud: d.env.AuthJwksAud, Iss: d.env.AuthJwksIss, Log: d.log, URL: d.env.AuthJwksUrl}).Handler,
				usermiddleware.NewMiddleware(usermiddleware.MiddlewareConfig{Log: d.log, Use: d.sto.User()}).Handler,
			},
		})
	}

	return srv
}

func (d *Daemon) Worker() *worker.Worker {
	var whn *workerhandler.Handler
	{
		whn = workerhandler.New(workerhandler.Config{
			Env: d.env,
			Loc: d.loc,
			Log: d.log,
			Sto: d.sto,
		})
	}

	var wrk *worker.Worker
	{
		wrk = worker.New(worker.Config{
			Han: whn.Hand(),
			Log: d.log,
			Res: d.res,
		})
	}

	return wrk
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
