package emitter

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/emitter/claimemitter"
	"github.com/uvio-network/apiserver/pkg/emitter/reputationemitter"
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
	cla claimemitter.Interface
	rep reputationemitter.Interface
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

	return &Emitter{
		cla: claimemitter.NewRescue(claimemitter.RescueConfig{Log: c.Log, Res: c.Res}),
		rep: reputationemitter.NewRescue(reputationemitter.RescueConfig{Log: c.Log, Res: c.Res}),
		use: useremitter.NewRescue(useremitter.RescueConfig{Log: c.Log, Res: c.Res}),
		uvx: uvxemitter.NewRescue(uvxemitter.RescueConfig{Log: c.Log, Res: c.Res}),
	}
}

func (e *Emitter) Claim() claimemitter.Interface {
	return e.cla
}

func (e *Emitter) Reputation() reputationemitter.Interface {
	return e.rep
}

func (e *Emitter) User() useremitter.Interface {
	return e.use
}

func (e *Emitter) UVX() uvxemitter.Interface {
	return e.uvx
}
