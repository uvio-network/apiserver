package fakeit

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
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
	var nam string
	{
		nam = fak.Username()
	}

	var obj *user.CreateI
	{
		obj = &user.CreateI{
			Object: []*user.CreateI_Object{
				{
					Public: &user.CreateI_Object_Public{
						Image: fmt.Sprintf("https://picsum.photos/seed/%s/48/48", nam),
						Name:  nam,
					},
				},
			},
		}
	}

	return obj
}
