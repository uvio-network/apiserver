package fakeit

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/brianvoe/gofakeit/v7/source"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/spf13/cobra"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/daemon"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/uvio-network/apiserver/pkg/server"
	"github.com/uvio-network/apiserver/pkg/worker"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, args []string) error {
	var err error

	var env envvar.Env
	{
		env = envvar.Load(envvar.Fake)
	}

	// --------------------------------------------------------------------- //

	var srv *server.Server
	var wrk *worker.Worker
	{
		srv, wrk, err = daemon.Create(env)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		go srv.Daemon()
		go wrk.Daemon()
	}

	// --------------------------------------------------------------------- //

	var key jwk.Key
	var set jwk.Set
	{
		key, set, err = newKey()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		go srvJwk(set)
	}

	// --------------------------------------------------------------------- //

	var cli Client
	{
		cli = NewClient(env)
	}

	var fak *gofakeit.Faker
	{
		fak = gofakeit.NewFaker(source.NewCrypto(), true)
	}

	// --------------------------------------------------------------------- //

	var use []*user.CreateO
	{
		use, err = r.createUser(cli, key, fak)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	fmt.Printf("%#v\n", use)

	// TODO select some users
	// TODO create some posts

	// TODO select some users
	// TODO create some votes

	return nil
}
