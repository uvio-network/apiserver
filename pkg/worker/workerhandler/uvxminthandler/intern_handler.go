package uvxminthandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type ExternHandlerConfig struct {
	Log logger.Interface
	Sto storage.Interface
}

type InternHandler struct {
	log logger.Interface
	sto storage.Interface
}

func NewExternHandler(c ExternHandlerConfig) *InternHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &InternHandler{
		log: c.Log,
		sto: c.Sto,
	}
}
