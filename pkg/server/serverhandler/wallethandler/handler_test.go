package wallethandler

import (
	"context"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
)

func tesCtx() context.Context {
	var str string
	fuzz.New().NilChance(0.5).Fuzz(&str)
	return userid.NewContext(context.Background(), objectid.ID(str))
}

func tesHan() wallet.API {
	return &wrapper{
		han: NewHandler(HandlerConfig{
			Log: logger.Fake(),
			Rec: reconciler.Fake(),
			Sto: storage.Fake(),
		}),
	}
}
