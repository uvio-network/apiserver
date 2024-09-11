package storage

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
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
	use userstorage.Interface
	vot votestorage.Interface
	wal walletstorage.Interface
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
			pos: poststorage.NewRedigo(poststorage.RedigoConfig{
				Log: c.Log,
				Red: c.Red,
			}),
			use: userstorage.NewRedis(userstorage.RedisConfig{
				Log: c.Log,
				Red: c.Red,
			}),
			vot: votestorage.NewRedis(votestorage.RedisConfig{
				Log: c.Log,
				Red: c.Red,
			}),
			wal: walletstorage.NewRedis(walletstorage.RedisConfig{
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

func (s *Storage) User() userstorage.Interface {
	return s.use
}

func (s *Storage) Vote() votestorage.Interface {
	return s.vot
}

func (s *Storage) Wallet() walletstorage.Interface {
	return s.wal
}
