package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/cmd/daemon"
	"github.com/uvio-network/apiserver/cmd/version"
	"github.com/xh3b4sd/tracer"
)

var (
	use = "apiserver"
	sho = "Golang based RPC apiserver."
	lon = "Golang based RPC apiserver."
)

func New() (*cobra.Command, error) {
	var err error

	// --------------------------------------------------------------------- //

	var cmdDae *cobra.Command
	{
		c := daemon.Config{}

		cmdDae, err = daemon.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var cmdVer *cobra.Command
	{
		c := version.Config{}

		cmdVer, err = version.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
			// We slience errors because we do not want to see spf13/cobra printing.
			// The errors returned by the commands will be propagated to the main.go
			// anyway, where we have custom error printing for the command line
			// tool.
			SilenceErrors: true,
			SilenceUsage:  true,
		}
	}

	{
		c.SetHelpCommand(&cobra.Command{Hidden: true})
	}

	{
		c.AddCommand(cmdDae)
		c.AddCommand(cmdVer)
	}

	return c, nil
}
