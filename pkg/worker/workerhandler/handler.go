package workerhandler

import (
	"fmt"

	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Loc locker.Interface
	Log logger.Interface
}

type Handler struct {
	han []Interface
}

func New(c Config) *Handler {
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	var han []Interface

	// {
	// 	han = append(han, descriptiondeletehandler.NewCustomHandler(descriptiondeletehandler.CustomHandlerConfig{
	// 		Des: c.Sto.Desc(),
	// 		Log: c.Log,
	// 	}))
	// }

	var h *Handler
	{
		h = &Handler{
			han: han,
		}
	}

	return h
}

func (h *Handler) Hand() []Interface {
	return h.han
}
