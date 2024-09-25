package reconciler

import (
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return New(Config{
		Log: logger.Fake(),
		Red: redigo.Fake(),
		Sto: storage.Fake(),
	})
}
