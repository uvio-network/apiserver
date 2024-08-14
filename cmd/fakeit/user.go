package fakeit

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createUser(cli Client, key jwk.Key, fak *gofakeit.Faker) (*user.SearchO, error) {
	var err error

	var ids []string
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

		for _, x := range out.Object {
			ids = append(ids, x.Intern.Id)
		}
	}

	var inp *user.SearchI
	{
		inp = &user.SearchI{}
	}

	for _, x := range ids {
		inp.Object = append(inp.Object, &user.SearchI_Object{
			Intern: &user.SearchI_Object_Intern{
				Id: x,
			},
		})
	}

	var out *user.SearchO
	{
		out, err = cli.User.Search(context.Background(), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return out, nil
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

func useObj(use *user.SearchO, uid string) *user.SearchO_Object {
	for _, x := range use.Object {
		if x.Intern.Id == uid {
			return x
		}
	}

	return nil
}
