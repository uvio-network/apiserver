package postreconciler

import (
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
)

func Fake() Interface {
	return NewRedigo(RedigoConfig{
		Log: logger.Fake(),
		Sto: storage.Fake(),
	})
}
