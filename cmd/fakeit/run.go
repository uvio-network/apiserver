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
	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/uvio-network/apiserver/pkg/daemon"
	"github.com/uvio-network/apiserver/pkg/envvar"
	"github.com/xh3b4sd/tracer"
)

type run struct {
	cli Client
	dae daemon.Interface
	fak *gofakeit.Faker
}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	// --------------------------------------------------------------------- //

	{
		r.dae = daemon.New(envvar.Load("fake"))
	}

	{
		go r.dae.Server().Daemon()
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
		go srvJwk(r.dae.Env(), set)
	}

	// --------------------------------------------------------------------- //

	{
		r.cli = NewClient(r.dae.Env())
	}

	{
		r.fak = gofakeit.NewFaker(source.NewCrypto(), true)
	}

	// --------------------------------------------------------------------- //

	var use *user.SearchO
	{
		use, err = r.createUser(key)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var wal *wallet.SearchO
	{
		wal, err = r.createWallet(key, use)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var pro *post.SearchO
	{
		pro, err = r.createPropose(key, use)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var vot *vote.SearchO
	{
		vot, err = r.createVote(key, use, pro)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var res *post.SearchO
	{
		res, err = r.createResolve(wal, pro)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var com *post.SearchO
	{
		com, err = r.createComment(key, use, pro, vot)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Printf("Generated %d fake users.\n", len(use.Object))
		fmt.Printf("Generated %d fake wallets.\n", len(wal.Object))
		fmt.Printf("Generated %d fake claims of lifecycle phase %q.\n", len(pro.Object), "propose")
		fmt.Printf("Generated %d fake claims of lifecycle phase %q.\n", len(res.Object), "resolve")
		fmt.Printf("Generated %d fake votes.\n", len(vot.Object))
		fmt.Printf("Generated %d fake comments.\n", len(com.Object))
	}

	return nil
}
