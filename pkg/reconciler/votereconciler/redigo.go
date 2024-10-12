package votereconciler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedigoConfig struct {
	Log logger.Interface
	Red redigo.Interface
	Sto storage.Interface
}

type Redigo struct {
	log logger.Interface
	red redigo.Interface
	sto storage.Interface
}

func NewRedigo(c RedigoConfig) *Redigo {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	return &Redigo{
		log: c.Log,
		red: c.Red,
		sto: c.Sto,
	}
}

func votOwnCla(uid objectid.ID, cid objectid.ID) string {
	return fmt.Sprintf(storageformat.VoteOwnerClaim, uid, cid)
}
