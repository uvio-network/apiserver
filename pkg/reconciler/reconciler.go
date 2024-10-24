package reconciler

import (
	"fmt"

	"github.com/uvio-network/apiserver/pkg/reconciler/postreconciler"
	"github.com/uvio-network/apiserver/pkg/reconciler/votereconciler"
	"github.com/uvio-network/apiserver/pkg/reconciler/walletreconciler"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/redigo"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
	Red redigo.Interface
	Sto storage.Interface
}

type Reconciler struct {
	pos postreconciler.Interface
	vot votereconciler.Interface
	wal walletreconciler.Interface
}

func New(c Config) *Reconciler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Red == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Red must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	var r *Reconciler
	{
		r = &Reconciler{
			pos: postreconciler.NewRedigo(postreconciler.RedigoConfig{
				Log: c.Log,
				Sto: c.Sto,
			}),
			vot: votereconciler.NewRedigo(votereconciler.RedigoConfig{
				Log: c.Log,
				Red: c.Red,
				Sto: c.Sto,
			}),
			wal: walletreconciler.NewRedigo(walletreconciler.RedigoConfig{
				Log: c.Log,
				Sto: c.Sto,
			}),
		}
	}

	return r
}

func (r *Reconciler) Post() postreconciler.Interface {
	return r.pos
}

func (r *Reconciler) Vote() votereconciler.Interface {
	return r.vot
}

func (r *Reconciler) Wallet() walletreconciler.Interface {
	return r.wal
}
