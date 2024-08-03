package votehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/validator"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Sto storage.Interface
	Val validator.Interface
}

type Handler struct {
	log logger.Interface
	sto storage.Interface
	val validator.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}
	if c.Val == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Val must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		val: c.Val,
		sto: c.Sto,
	}
}
