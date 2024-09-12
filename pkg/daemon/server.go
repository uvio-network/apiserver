package daemon

import (
	"github.com/gorilla/mux"
	"github.com/twitchtv/twirp"
	"github.com/uvio-network/apiserver/pkg/server"
	"github.com/uvio-network/apiserver/pkg/server/interceptor/failedinterceptor"
	"github.com/uvio-network/apiserver/pkg/server/middleware/authmiddleware"
	"github.com/uvio-network/apiserver/pkg/server/middleware/corsmiddleware"
	"github.com/uvio-network/apiserver/pkg/server/middleware/usermiddleware"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler"
)

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

	return server.New(server.Config{
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