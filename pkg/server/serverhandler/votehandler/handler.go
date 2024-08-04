package votehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Rec reconciler.Interface
	Sto storage.Interface
}

type Handler struct {
	log logger.Interface
	rec reconciler.Interface
	sto storage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rec == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rec must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		rec: c.Rec,
		sto: c.Sto,
	}
}
