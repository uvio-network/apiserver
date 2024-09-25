package postreconciler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type RedigoConfig struct {
	Log logger.Interface
	Sto storage.Interface
}

type Redigo struct {
	log logger.Interface
	sto storage.Interface
}

func NewRedigo(c RedigoConfig) *Redigo {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &Redigo{
		log: c.Log,
		sto: c.Sto,
	}
}
