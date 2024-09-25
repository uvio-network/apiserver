package fakeit

import (
	"context"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createWallet(key jwk.Key, use *user.SearchO) (*wallet.SearchO, error) {
	var err error

	all := map[string]struct{}{}

	for i := 0; i < len(use.Object); i++ {
		var nam string
		{
			nam = use.Object[i].Public.Name
		}

		{
			all[nam] = struct{}{}
		}

		var inp *wallet.CreateI
		{
			inp = r.randomWallet()
		}

		var ctx context.Context
		{
			ctx = newCtx(key, nam)
		}

		{
			_, err = r.cli.Wallet.Create(ctx, inp)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}
	}

	// Do it again randomized, so that some users have multiple wallets.

	for i := 0; i < len(use.Object); i++ {
		{
			r.fak.ShuffleAnySlice(use.Object)
		}

		var nam string
		{
			nam = use.Object[0].Public.Name
		}

		{
			all[nam] = struct{}{}
		}

		var inp *wallet.CreateI
		{
			inp = r.randomWallet()
		}

		var ctx context.Context
		{
			ctx = newCtx(key, nam)
		}

		{
			_, err = r.cli.Wallet.Create(ctx, inp)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}
	}

	var out *wallet.SearchO
	{
		out = &wallet.SearchO{}
	}

	for k := range all {
		var inp *wallet.SearchI
		{
			inp = &wallet.SearchI{
				Object: []*wallet.SearchI_Object{
					{
						Intern: &wallet.SearchI_Object_Intern{
							Owner: "self",
						},
					},
				},
			}
		}

		res, err := r.cli.Wallet.Search(newCtx(key, k), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}

		{
			out.Object = append(out.Object, res.Object...)
		}
	}

	return out, nil
}

func (r *run) randomWallet() *wallet.CreateI {
	var kin []string
	{
		kin = []string{
			"embedded",
			"injected",
		}
	}

	var obj *wallet.CreateI
	{
		obj = &wallet.CreateI{
			Object: []*wallet.CreateI_Object{
				{
					Public: &wallet.CreateI_Object_Public{
						Address:     r.fak.HexUint(160),
						Description: r.fak.BookTitle(),
						Kind:        r.fak.RandomString(kin),
					},
				},
			},
		}
	}

	return obj
}
