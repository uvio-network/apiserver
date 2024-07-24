package postdeletehandler

import (
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type CustomHandlerConfig struct {
	Log logger.Interface
}

type CustomHandler struct {
	log logger.Interface
}

func NewCustomHandler(c CustomHandlerConfig) *CustomHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &CustomHandler{
		log: c.Log,
	}
}
