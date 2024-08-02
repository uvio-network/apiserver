package validator

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/validator/postvalidator"
	"github.com/uvio-network/apiserver/pkg/validator/votevalidator"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
	Sto storage.Interface
}

type Validator struct {
	pos postvalidator.Interface
	vot votevalidator.Interface
}

func New(c Config) *Validator {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	var v *Validator
	{
		v = &Validator{
			pos: postvalidator.NewRedis(postvalidator.RedisConfig{
				Log: c.Log,
				Sto: c.Sto,
			}),
			vot: votevalidator.NewRedis(votevalidator.RedisConfig{
				Log: c.Log,
				Sto: c.Sto,
			}),
		}
	}

	return v
}

func (v *Validator) Post() postvalidator.Interface {
	return v.pos
}

func (v *Validator) Vote() votevalidator.Interface {
	return v.vot
}
