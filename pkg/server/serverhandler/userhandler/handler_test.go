package userhandler

import (
	"context"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/emitter"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/objectid"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() user.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Emi: emitter.Fake(),
			Log: logger.Fake(),
			Sto: storage.Fake(),
		}),
	}
}
