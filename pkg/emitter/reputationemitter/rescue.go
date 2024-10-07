package reputationemitter

import (
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/rescue"
	"github.com/xh3b4sd/tracer"
)

type RescueConfig struct {
	Log logger.Interface
	Res rescue.Interface
}

type Rescue struct {
	log logger.Interface
	res rescue.Interface
}

func NewRescue(c RescueConfig) *Rescue {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	return &Rescue{
		log: c.Log,
		res: c.Res,
	}
}
