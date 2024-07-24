package storage

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
	Red redigo.Interface
}

type Storage struct {
	pos poststorage.Interface
}

func New(c Config) *Storage {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	var s *Storage
	{
		s = &Storage{
			pos: poststorage.NewRedis(poststorage.RedisConfig{
				Log: c.Log,
				Red: c.Red,
			}),
		}
	}

	return s
}

func (s *Storage) Post() poststorage.Interface {
	return s.pos
}
