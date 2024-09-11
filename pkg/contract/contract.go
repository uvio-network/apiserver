package contract

import (
	"fmt"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
}

type Contract struct {
	log logger.Interface
}

func New(c Config) *Contract {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Contract{
		log: c.Log,
	}
}

// func (e *Contract) User() useremitter.Interface {
// 	return e.use
// }
