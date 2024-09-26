package rescue

import (
	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/pkg/daemon"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	var dae daemon.Interface
	{
		dae = daemon.New(envvar.Load("local"))
	}

	// --------------------------------------------------------------------- //

	//
	//     ./apiserver rescue user.uvio.network/object 1727370815215919
	//

	var key string
	{
		key = arg[0]
	}

	var val string
	{
		val = arg[1]
	}

	// --------------------------------------------------------------------- //

	var lis []*task.Task
	{
		lis, err = dae.Res().Lister(&task.Task{Meta: &task.Meta{key: val}})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	if len(lis) != 1 {
		return tracer.Maskf(runtime.ExecutionFailedError, "expected to find exactly one task, found %d", len(lis))
	}

	{
		err = dae.Res().Cycles(lis[0])
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
