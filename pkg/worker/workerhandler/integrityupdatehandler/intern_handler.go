package integrityupdatehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/contract"
	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type InternHandlerConfig struct {
	Con contract.Interface
	Emi emitter.Interface
	Log logger.Interface
	Sto storage.Interface
}

type InternHandler struct {
	con contract.Interface
	emi emitter.Interface
	log logger.Interface
	sto storage.Interface
}

func NewInternHandler(c InternHandlerConfig) *InternHandler {
	if c.Con == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Con must not be empty", c)))
	}
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &InternHandler{
		con: c.Con,
		emi: c.Emi,
		log: c.Log,
		sto: c.Sto,
	}
}
