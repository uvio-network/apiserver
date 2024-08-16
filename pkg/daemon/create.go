package daemon

import (
	"net"
	"time"

	"github.com/gorilla/mux"
	"github.com/twitchtv/twirp"
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

func Create(env envvar.Env) (*server.Server, *worker.Worker, error) {
	var err error

	var log logger.Interface
	{
		log = logger.Default()
	}

	var lis net.Listener
	{
		lis, err = net.Listen("tcp", net.JoinHostPort(env.HttpHost, env.HttpPort))
		if err != nil {
			return nil, nil, tracer.Mask(err)
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

	var shn *serverhandler.Handler
	{
		shn = serverhandler.New(serverhandler.Config{
			Loc: loc,
			Log: log,
			Rec: rec,
			Sto: sto,
		})
	}

	var srv *server.Server
	{
		srv = server.New(server.Config{
			Han: shn.Hand(),
			Int: []twirp.Interceptor{
				failedinterceptor.NewInterceptor(failedinterceptor.InterceptorConfig{Log: log}).Interceptor,
			},
			Lis: lis,
			Log: log,
			Mid: []mux.MiddlewareFunc{
				corsmiddleware.NewMiddleware(corsmiddleware.MiddlewareConfig{Log: log}).Handler,
				authmiddleware.NewMiddleware(authmiddleware.MiddlewareConfig{Aud: env.AuthJwksAud, Iss: env.AuthJwksIss, Log: log, URL: env.AuthJwksUrl}).Handler,
				usermiddleware.NewMiddleware(usermiddleware.MiddlewareConfig{Log: log, Use: sto.User()}).Handler,
			},
		})
	}

	// --------------------------------------------------------------------- //

	var whn *workerhandler.Handler
	{
		whn = workerhandler.New(workerhandler.Config{
			Loc: loc,
			Log: log,
			Sto: sto,
		})
	}

	var wrk *worker.Worker
	{
		wrk = worker.New(worker.Config{
			Han: whn.Hand(),
			Log: log,
			Res: res,
		})
	}

	return srv, wrk, nil
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
