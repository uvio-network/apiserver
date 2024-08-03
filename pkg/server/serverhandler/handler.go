package serverhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/server/serverhandler/posthandler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/userhandler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/votehandler"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/validator"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Loc locker.Interface
	Log logger.Interface
	Sto storage.Interface
	Val validator.Interface
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
			Sto: c.Sto,
			Val: c.Val,
		}))
	}

	{
		han = append(han, userhandler.NewHandler(userhandler.HandlerConfig{
			Log: c.Log,
			Sto: c.Sto,
		}))
	}

	{
		han = append(han, votehandler.NewHandler(votehandler.HandlerConfig{
			Log: c.Log,
			Sto: c.Sto,
			Val: c.Val,
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
