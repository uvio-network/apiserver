package redigo

import (
	"github.com/spf13/cobra"
)

const (
	use = "redigo"
	sho = "Manipulate redis data."
	lon = "Manipulate redis data."
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
