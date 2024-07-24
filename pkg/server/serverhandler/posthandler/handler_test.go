package posthandler

import (
	"context"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() post.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
		}),
	}
}
