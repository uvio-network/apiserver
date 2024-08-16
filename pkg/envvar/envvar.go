package envvar

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	Fake  = ".env.fake"
	Local = ".env.local"
)

type Env struct {
	AuthJwksAud      string `split_words:"true" required:"true"`
	AuthJwksIss      string `split_words:"true" required:"true"`
	AuthJwksUrl      string `split_words:"true" required:"true"`
	ChainId          string `split_words:"true" required:"true"`
	ChainRpcEndpoint string `split_words:"true" required:"true"`
	HttpHost         string `split_words:"true" default:"127.0.0.1"`
	HttpPort         string `split_words:"true" default:"7777"`
	MarketsAddress   string `split_words:"true" required:"true"`
}

func Load(pat string) Env {
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
