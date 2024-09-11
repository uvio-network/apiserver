package emitter

import (
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
)

func Fake() Interface {
	return New(Config{
		Log: logger.Fake(),
		Res: rescue.Fake(),
	})
}
