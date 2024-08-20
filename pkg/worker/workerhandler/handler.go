package workerhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/claimresolvehandler"
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

	var han []Interface

	{
		han = append(han, claimresolvehandler.NewSystemHandler(claimresolvehandler.SystemHandlerConfig{
			Cid: c.Env.ChainId,
			Pk:  c.Env.PrivateKey,
			Log: c.Log,
			Rpc: c.Env.ChainRpcEndpoint,
			Sto: c.Sto,
			Cas: claimresolvehandler.ContractAddresses{
				Markets:    c.Env.MarketsAddress,
				Randomizer: c.Env.RandomizerAddress,
			},
		}))
	}

	var h *Handler
	{
		h = &Handler{
			han: han,
		}
	}

	return h
}

func (h *Handler) Hand() []Interface {
	return h.han
}
