package workerhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/contract"
	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/balanceupdatehandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/claimexpiryhandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/integrityupdatehandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/resolvecreatehandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/usercreatehandler"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler/uvxminthandler"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Con contract.Interface
	Emi emitter.Interface
	Env envvar.Env
	Loc locker.Interface
	Log logger.Interface
	Rec reconciler.Interface
	Sam *sample.Sample
	Sto storage.Interface
}

type Handler struct {
	han []Interface
}

func New(c Config) *Handler {
	if c.Con == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Con must not be empty", c)))
	}
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rec == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rec must not be empty", c)))
	}
	if c.Sam == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sam must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	var han []Interface

	{
		han = append(han, balanceupdatehandler.NewInternHandler(balanceupdatehandler.InternHandlerConfig{
			Con: c.Con,
			Emi: c.Emi,
			Log: c.Log,
			Rec: c.Rec,
			Sto: c.Sto,
		}))

		han = append(han, claimexpiryhandler.NewInternHandler(claimexpiryhandler.InternHandlerConfig{
			Con: c.Con,
			Emi: c.Emi,
			Log: c.Log,
			Sto: c.Sto,
		}))

		han = append(han, integrityupdatehandler.NewInternHandler(integrityupdatehandler.InternHandlerConfig{
			Con: c.Con,
			Log: c.Log,
			Sto: c.Sto,
		}))

		han = append(han, resolvecreatehandler.NewInternHandler(resolvecreatehandler.InternHandlerConfig{
			Con: c.Con,
			Log: c.Log,
			Rec: c.Rec,
			Sam: c.Sam,
			Sto: c.Sto,
		}))

		han = append(han, usercreatehandler.NewExternHandler(usercreatehandler.ExternHandlerConfig{
			Emi: c.Emi,
			Log: c.Log,
			Sto: c.Sto,
		}))

		han = append(han, uvxminthandler.NewInternHandler(uvxminthandler.InternHandlerConfig{
			Con: c.Con,
			Log: c.Log,
			Sto: c.Sto,
			UVX: c.Env.ChainUvxContract,
		}))
	}

	return &Handler{
		han: han,
	}
}

func (h *Handler) Hand() []Interface {
	return h.han
}
