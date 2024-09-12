package daemon

import (
	"github.com/uvio-network/apiserver/pkg/contract"
	"github.com/uvio-network/apiserver/pkg/worker"
	"github.com/uvio-network/apiserver/pkg/worker/workerhandler"
)

func (d *Daemon) Worker() *worker.Worker {
	var con contract.Interface
	{
		con = contract.New(contract.Config{
			Key: d.env.SignerPrivateKey,
			Log: d.log,
			RPC: d.env.ChainRpcEndpoint,
			UVX: d.env.ChainUvxContract,
		})
	}

	var whn *workerhandler.Handler
	{
		whn = workerhandler.New(workerhandler.Config{
			Con: con,
			Emi: d.emi,
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
