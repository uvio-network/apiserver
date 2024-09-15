package walletstorage

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
)

func Fake() Interface {
	return NewRedigo(RedigoConfig{
		Log: logger.Fake(),
		Red: redigo.Fake(),
	})
}
