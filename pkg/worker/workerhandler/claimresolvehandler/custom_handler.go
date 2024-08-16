package claimresolvehandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Log logger.Interface
	Sto storage.Interface
}

type SystemHandler struct {
	log logger.Interface
	sto storage.Interface
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &SystemHandler{
		log: c.Log,
		sto: c.Sto,
	}
}
