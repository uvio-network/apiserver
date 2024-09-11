package daemon

import (
	"github.com/uvio-network/apiserver/pkg/worker"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler"
)

func (d *Daemon) Worker() *worker.Worker {
	var whn *workerhandler.Handler
	{
		whn = workerhandler.New(workerhandler.Config{
			Env: d.env,
			Loc: d.loc,
			Log: d.log,
			Sto: d.sto,
		})
	}

	return worker.New(worker.Config{
		Han: whn.Hand(),
		Log: d.log,
		Res: d.res,
	})
}
