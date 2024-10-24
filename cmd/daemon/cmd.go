package daemon

import (
	"github.com/spf13/cobra"
)

const (
	use = "daemon"
	sho = "Execute the long running process exposing RPC server handlers."
	lon = "Execute the long running process exposing RPC server handlers."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var f *flag
	{
		f = &flag{}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{flag: f}).runE,
		}
	}

	{
		f.Init(c)
	}

	return c, nil
}
