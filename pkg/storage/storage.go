package storage

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
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
	not notestorage.Interface
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

	return &Storage{
		not: notestorage.NewRedigo(notestorage.RedigoConfig{
			Log: c.Log,
			Red: c.Red,
		}),
		pos: poststorage.NewRedigo(poststorage.RedigoConfig{
			Log: c.Log,
			Red: c.Red,
		}),
		use: userstorage.NewRedigo(userstorage.RedigoConfig{
			Log: c.Log,
			Red: c.Red,
		}),
		vot: votestorage.NewRedigo(votestorage.RedigoConfig{
			Log: c.Log,
			Red: c.Red,
		}),
		wal: walletstorage.NewRedigo(walletstorage.RedigoConfig{
			Log: c.Log,
			Red: c.Red,
		}),
	}
}

func (s *Storage) Note() notestorage.Interface {
	return s.not
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
