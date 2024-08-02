package votehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/uvio-network/apiserver/pkg/validator/votevalidator"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Sto votestorage.Interface
	Val votevalidator.Interface
}

type Handler struct {
	log logger.Interface
	sto votestorage.Interface
	val votevalidator.Interface
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
