package posthandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type HandlerConfig struct {
	Log logger.Interface
	Pos poststorage.Interface
}

type Handler struct {
	log logger.Interface
	pos poststorage.Interface
}

func NewHandler(c HandlerConfig) *Handler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Pos == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Pos must not be empty", c)))
	}

	return &Handler{
		log: c.Log,
		pos: c.Pos,
	}
}
