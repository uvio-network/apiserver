package signer

import (
	"github.com/spf13/cobra"
)

const (
	use = "signer"
	sho = "Create a new wallet."
	lon = "Create a new wallet."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
		}
	}

	return c, nil
}
