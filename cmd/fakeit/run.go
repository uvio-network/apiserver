package fakeit

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/brianvoe/gofakeit/v7/source"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/spf13/cobra"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/daemon"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	var env envvar.Env
	{
		env = envvar.Load("fake")
	}

	// --------------------------------------------------------------------- //

	var dae *daemon.Daemon
	{
		dae = daemon.New(env)
	}

	{
		go dae.Server().Daemon()
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

	var use *user.SearchO
	{
		use, err = r.createUser(cli, key, fak)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var cla *post.SearchO
	{
		cla, err = r.createClaim(cli, key, fak, use)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var vot *vote.SearchO
	{
		vot, err = r.createVote(cli, key, fak, use, cla)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var com *post.SearchO
	{
		com, err = r.createComment(cli, key, fak, use, cla, vot)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Printf("Generated %d fake users.\n", len(use.Object))
		fmt.Printf("Generated %d fake claims.\n", len(cla.Object))
		fmt.Printf("Generated %d fake votes.\n", len(vot.Object))
		fmt.Printf("Generated %d fake comments.\n", len(com.Object))
	}

	return nil
}
