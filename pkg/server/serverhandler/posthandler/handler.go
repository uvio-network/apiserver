package posthandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/validator/postvalidator"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Sto poststorage.Interface
	Val postvalidator.Interface
}

type Handler struct {
	log logger.Interface
	sto poststorage.Interface
	val postvalidator.Interface
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
		sto: c.Sto,
		val: c.Val,
	}
}
