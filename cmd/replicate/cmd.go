package replicate

import (
	"github.com/spf13/cobra"
)

const (
	use = "replicate"
	sho = "Replicate redis data."
	lon = "Replicate redis data."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{}).runE,
		}
	}

	return c, nil
}
