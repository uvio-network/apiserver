package walletstorage

import (
	"encoding/json"
	"fmt"

	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type RedigoConfig struct {
	Log logger.Interface
	Red redigo.Interface
}

type Redigo struct {
	log logger.Interface
	red redigo.Interface
}

func NewRedigo(c RedigoConfig) *Redigo {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}

	return &Redigo{
		log: c.Log,
		red: c.Red,
	}
}

func walAdd(str string) string {
	return fmt.Sprintf(storageformat.WalletAddress, str)
}

func walObj(oid objectid.ID) string {
	return fmt.Sprintf(storageformat.WalletObject, oid)
}

func walOwn(oid objectid.ID) string {
	return fmt.Sprintf(storageformat.WalletOwner, oid)
}

func musStr(obj *Object) string {
	byt, err := json.Marshal(obj)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
