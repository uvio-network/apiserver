package serverhandler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/notehandler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/posthandler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/userhandler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/votehandler"
	"github.com/uvio-network/apiserver/pkg/server/serverhandler/wallethandler"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/locker"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Emi emitter.Interface
	Loc locker.Interface
	Log logger.Interface
	Rec reconciler.Interface
	Sto storage.Interface
}

type Handler struct {
	han []Interface
}

func New(c Config) *Handler {
	if c.Emi == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Emi must not be empty", c)))
	}
	if c.Loc == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Loc must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	var han []Interface

	{
		han = append(han, notehandler.NewHandler(notehandler.HandlerConfig{
			Log: c.Log,
			Rec: c.Rec,
			Sto: c.Sto,
		}))

		han = append(han, posthandler.NewHandler(posthandler.HandlerConfig{
			Log: c.Log,
			Rec: c.Rec,
			Sto: c.Sto,
		}))

		han = append(han, userhandler.NewHandler(userhandler.HandlerConfig{
			Emi: c.Emi,
			Log: c.Log,
			Sto: c.Sto,
		}))

		han = append(han, votehandler.NewHandler(votehandler.HandlerConfig{
			Log: c.Log,
			Rec: c.Rec,
			Sto: c.Sto,
		}))

		han = append(han, wallethandler.NewHandler(wallethandler.HandlerConfig{
			Log: c.Log,
			Rec: c.Rec,
			Sto: c.Sto,
		}))
	}

	return &Handler{
		han: han,
	}
}

func (h *Handler) Hand() []Interface {
	return h.han
}
