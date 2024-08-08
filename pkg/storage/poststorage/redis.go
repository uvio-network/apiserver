package poststorage

import (
	"encoding/json"
	"fmt"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedisConfig struct {
	Log logger.Interface
	Red redigo.Interface
}

type Redis struct {
	log logger.Interface
	red redigo.Interface
}

func NewRedis(c RedisConfig) *Redis {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Redis{
		log: c.Log,
		red: c.Red,
	}
}

func posCom(oid objectid.ID) string {
	return fmt.Sprintf(storageformat.PostComment, oid)
}

func posUseCom(uid objectid.ID, pid objectid.ID) string {
	return fmt.Sprintf(storageformat.PostUserComment, uid, pid)
}

func posLab(str string) string {
	return fmt.Sprintf(storageformat.PostLabel, str)
}

func posObj(oid objectid.ID) string {
	return fmt.Sprintf(storageformat.PostObject, oid)
}

func posTre(oid objectid.ID) string {
	return fmt.Sprintf(storageformat.PostTree, oid)
}

func posOwn(oid objectid.ID) string {
	return fmt.Sprintf(storageformat.PostOwner, oid)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
