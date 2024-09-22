package walletreconciler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Sto storage.Interface
}

type Redis struct {
	log logger.Interface
	sto storage.Interface
}

// TODO rename reconciler
func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		sto: c.Sto,
	}
}
