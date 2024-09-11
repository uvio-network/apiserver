package uvxminthandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/contract"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type InternHandlerConfig struct {
	Con *contract.Contract
	Log logger.Interface
	Sto storage.Interface
}

type InternHandler struct {
	con *contract.Contract
	log logger.Interface
	sto storage.Interface
}

func NewInternHandler(c InternHandlerConfig) *InternHandler {
	if c.Con == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Con must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &InternHandler{
		con: c.Con,
		log: c.Log,
		sto: c.Sto,
	}
}
