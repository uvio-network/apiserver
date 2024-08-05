package fakeit

import (
	"context"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createUser(cli Client, key jwk.Key, fak *gofakeit.Faker) ([]*user.CreateO, error) {
	var err error
	var lis []*user.CreateO

	for i := 0; i < 10; i++ {
		var inp *user.CreateI
		{
			inp = r.randomUser(fak)
		}

		var ctx context.Context
		{
			ctx = newCtx(key, inp.Object[0].Public.Name)
		}

		var out *user.CreateO
		{
			out, err = cli.User.Create(ctx, inp)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		lis = append(lis, out)
	}

	return lis, nil
}

func (r *run) randomUser(fak *gofakeit.Faker) *user.CreateI {
	var err error

	// We want to produce a fake profile picture for our fake user. Here we fetch
	// a random image URL from some image provider. For that to work we use a
	// custom HTTP client that is not following redirects, because the random
	// image URL we are looking for will be set in the location header of the HTTP
	// response below.
	var cli *http.Client
	{
		cli = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	var res *http.Response
	{
		res, err = cli.Get(fak.ImageURL(48, 48))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var obj *user.CreateI
	{
		obj = &user.CreateI{
			Object: []*user.CreateI_Object{
				{
					Public: &user.CreateI_Object_Public{
						Image: res.Header.Get("location"),
						Name:  fak.Username(),
					},
				},
			},
		}
	}

	return obj
}
