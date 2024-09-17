package daemon

import (
	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type flag struct {
	Env string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Env, "env", "local", "the environment file to load, e.g. local for env.local")
}

func (f *flag) Validate() error {
	if f.Env != "fake" && f.Env != "local" && f.Env != "sepolia" {
		return tracer.Maskf(runtime.ExecutionFailedError, "--env must be either local, fake or sepolia")
	}

	return nil
}
