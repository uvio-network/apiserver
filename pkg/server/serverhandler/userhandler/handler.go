package userhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Emi emitter.Interface
	Log logger.Interface
	Sto storage.Interface
}

type Handler struct {
	emi emitter.Interface
	log logger.Interface
	sto storage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &Handler{
		emi: c.Emi,
		log: c.Log,
		sto: c.Sto,
	}
}
