package envvar

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	pat = ".env.local"
)

type Env struct {
	HttpHost string `split_words:"true" default:"127.0.0.1"`
	HttpPort string `split_words:"true" default:"7777"`
}

func Load() Env {
	var err error

	var env Env

	for {
		{
			err = godotenv.Load(pat)
			if err != nil {
				fmt.Printf("could not load %s (%s)\n", pat, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		{
			err = envconfig.Process("APISERVER", &env)
			if err != nil {
				fmt.Printf("could not process envfile %s (%s)\n", pat, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		return env
	}
}
