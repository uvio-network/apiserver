package envvar

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	AuthJwksAud         string `split_words:"true" required:"true"`
	AuthJwksIss         string `split_words:"true" required:"true"`
	AuthJwksUrl         string `split_words:"true" required:"true"`
	ChainClaimsContract string `split_words:"true" required:"true"`
	ChainRpcEndpoint    string `split_words:"true" required:"true"`
	ChainUvxContract    string `split_words:"true" required:"true"`
	HttpHost            string `split_words:"true" required:"true"`
	HttpPort            string `split_words:"true" required:"true"`
	LogLevel            string `split_words:"true" default:"debug"`
	SignerPrivateKey    string `split_words:"true" required:"true"`
}

func Load(kin string) Env {
	var err error

	var env Env

	for {
		{
			err = godotenv.Load(fmt.Sprintf(".env.%s", kin))
			if err != nil {
				fmt.Printf("could not load %s (%s)\n", kin, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		{
			err = envconfig.Process("APISERVER", &env)
			if err != nil {
				fmt.Printf("could not process envfile %s (%s)\n", kin, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		return env
	}
}
