package daemon

import (
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/reconciler"
	"github.com/uvio-network/apiserver/pkg/server"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/uvio-network/apiserver/pkg/worker"
)

type Interface interface {
	Env() envvar.Env
	Rec() reconciler.Interface
	Sto() storage.Interface

	Server() *server.Server
	Worker() *worker.Worker
}
