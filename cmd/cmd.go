package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uvio-network/apiserver/cmd/daemon"
	"github.com/uvio-network/apiserver/cmd/fakeit"
	"github.com/uvio-network/apiserver/cmd/redigo"
	"github.com/uvio-network/apiserver/cmd/rescue"
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

	var dae *cobra.Command
	{
		c := daemon.Config{}

		dae, err = daemon.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var fak *cobra.Command
	{
		c := fakeit.Config{}

		fak, err = fakeit.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var red *cobra.Command
	{
		c := redigo.Config{}

		red, err = redigo.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var res *cobra.Command
	{
		c := rescue.Config{}

		res, err = rescue.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var ver *cobra.Command
	{
		c := version.Config{}

		ver, err = version.New(c)
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
		c.AddCommand(dae)
		c.AddCommand(fak)
		c.AddCommand(red)
		c.AddCommand(res)
		c.AddCommand(ver)
	}

	return c, nil
}
