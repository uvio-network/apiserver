package serverhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/server/serverhandler/posthandler"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Loc locker.Interface
	Log logger.Interface
	Sto storage.Interface
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

	{
		han = append(han, posthandler.NewHandler(posthandler.HandlerConfig{
			Log: c.Log,
			Pos: c.Sto.Post(),
		}))
	}

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
