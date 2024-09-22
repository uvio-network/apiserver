package daemon

import (
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/server"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/worker"
	"github.com/xh3b4sd/redigo"
)

type Interface interface {
	Env() envvar.Env
	Rec() reconciler.Interface
	Red() redigo.Interface
	Sto() storage.Interface

	Server() *server.Server
	Worker() *worker.Worker
}
