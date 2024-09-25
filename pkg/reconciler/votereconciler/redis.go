package votereconciler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Red redigo.Interface
	Sto storage.Interface
}

type Redis struct {
	log logger.Interface
	red redigo.Interface
	sto storage.Interface
}

// TODO rename reconciler
func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		red: c.Red,
		sto: c.Sto,
	}
}

func votOwnCla(uid objectid.ID, cid objectid.ID) string {
	return fmt.Sprintf(storageformat.VoteOwnerClaim, uid, cid)
}
