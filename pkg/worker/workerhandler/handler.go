package workerhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/contract"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/claimresolvehandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/usercreatehandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/uvxminthandler"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Env envvar.Env
	Loc locker.Interface
	Log logger.Interface
	Sto storage.Interface
}

type Handler struct {
	han []Interface
}

func New(c Config) *Handler {
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	var con *contract.Contract
	{
		con = contract.New(contract.Config{
			Key: c.Env.SignerPrivateKey,
			Log: c.Log,
			RPC: c.Env.ChainRpcEndpoint,
			UVX: c.Env.ChainUvxContract,
		})
	}

	var han []Interface

	{
		han = append(han, claimresolvehandler.NewInternHandler(claimresolvehandler.InternHandlerConfig{
			Log: c.Log,
			Sto: c.Sto,
		}))

		han = append(han, usercreatehandler.NewExternHandler(usercreatehandler.ExternHandlerConfig{
			Log: c.Log,
			Sto: c.Sto,
		}))

		han = append(han, uvxminthandler.NewInternHandler(uvxminthandler.InternHandlerConfig{
			Con: con,
			Log: c.Log,
			Sto: c.Sto,
		}))
	}

	return &Handler{
		han: han,
	}
}

func (h *Handler) Hand() []Interface {
	return h.han
}
