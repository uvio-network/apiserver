package createresolvehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/contract"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/sample"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type InternHandlerConfig struct {
	Con contract.Interface
	Log logger.Interface
	Rec reconciler.Interface
	Sam *sample.Sample
	Sto storage.Interface
}

type InternHandler struct {
	con contract.Interface
	log logger.Interface
	rec reconciler.Interface
	sam *sample.Sample
	sto storage.Interface
}

func NewInternHandler(c InternHandlerConfig) *InternHandler {
	if c.Con == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Con must not be empty", c)))
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

	return &InternHandler{
		con: c.Con,
		log: c.Log,
		rec: c.Rec,
		sam: c.Sam,
		sto: c.Sto,
	}
}
