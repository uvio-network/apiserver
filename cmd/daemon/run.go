package daemon

import (
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/server"
	"github.com/uvio-network/apiserver/pkg/server/interceptor/failedinterceptor"
	"github.com/uvio-network/apiserver/pkg/server/middleware/corsmiddleware"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler"
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

type run struct{}

func (r *run) runE(cmd *cobra.Command, args []string) error {
	var err error

	// --------------------------------------------------------------------- //

	var env envvar.Env
	{
		env = envvar.Load()
	}

	// --------------------------------------------------------------------- //

	var log logger.Interface
	{
		log = logger.Default()
	}

	var lis net.Listener
	{
		lis, err = net.Listen("tcp", net.JoinHostPort(env.HttpHost, env.HttpPort))
		if err != nil {
			return tracer.Mask(err)
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

	// --------------------------------------------------------------------- //

	var shn *serverhandler.Handler
	{
		shn = serverhandler.New(serverhandler.Config{
			Loc: loc,
			Log: log,
		})
	}

	// --------------------------------------------------------------------- //

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
			},
		})
	}

	{
		go srv.Daemon()
	}

	// --------------------------------------------------------------------- //

	var whn *workerhandler.Handler
	{
		whn = workerhandler.New(workerhandler.Config{
			Loc: loc,
			Log: log,
		})
	}

	// --------------------------------------------------------------------- //

	var wrk *worker.Worker
	{
		wrk = worker.New(worker.Config{
			Han: whn.Hand(),
			Log: log,
			Res: res,
		})
	}

	{
		go wrk.Daemon()
	}

	// --------------------------------------------------------------------- //

	var sig chan os.Signal
	{
		sig = make(chan os.Signal, 2)
	}

	{
		defer close(sig)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	}

	{
		<-sig
	}

	select {
	case <-time.After(10 * time.Second):
		// One SIGTERM gives the daemon some time to tear down gracefully.
	case <-sig:
		// Two SIGTERMs stop the immediatelly.
	}

	return nil
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

func splNum(str string) []int64 {
	var lis []int64

	for _, x := range strings.Split(str, ",") {
		lis = append(lis, musNum(x))
	}

	return lis
}

func splStr(str string) []string {
	return strings.Split(str, ",")
}

func musNum(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return num
}
