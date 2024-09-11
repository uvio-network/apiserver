package daemon

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/pkg/daemon"
	"github.com/uvio-network/apiserver/pkg/envvar"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, args []string) error {
	var env envvar.Env
	{
		env = envvar.Load(envvar.Local)
	}

	// --------------------------------------------------------------------- //

	var dae *daemon.Daemon
	{
		dae = daemon.New(env)
	}

	{
		go dae.Server().Daemon()
		go dae.Worker().Daemon()
	}

	// --------------------------------------------------------------------- //

	var sig chan os.Signal
	{
		sig = make(chan os.Signal, 2)
	}

	{
		defer close(sig)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	}

	{
		<-sig
	}

	select {
	case <-time.After(10 * time.Second):
		// One SIGTERM gives the daemon some time to tear down gracefully.
	case <-sig:
		// Two SIGTERMs stop the immediatelly.
	}

	return nil
}
