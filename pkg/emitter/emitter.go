package emitter

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/emitter/useremitter"
	"github.com/uvio-network/apiserver/pkg/emitter/uvxemitter"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
	Res rescue.Interface
}

type Emitter struct {
	use useremitter.Interface
	uvx uvxemitter.Interface
}

func New(c Config) *Emitter {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	var e *Emitter
	{
		e = &Emitter{
			use: useremitter.NewRescue(useremitter.RescueConfig{Log: c.Log, Res: c.Res}),
			uvx: uvxemitter.NewRescue(uvxemitter.RescueConfig{Log: c.Log, Res: c.Res}),
		}
	}

	return e
}

func (e *Emitter) User() useremitter.Interface {
	return e.use
}

func (e *Emitter) UVX() uvxemitter.Interface {
	return e.uvx
}
