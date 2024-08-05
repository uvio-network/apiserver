package fakeit

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createVote(cli Client, key jwk.Key, fak *gofakeit.Faker, use *user.SearchO, pos *post.SearchO) (*vote.SearchO, error) {
	var err error

	var ids []string
	for i := 0; i < 100; i++ {
		{
			fak.ShuffleAnySlice(use.Object)
			fak.ShuffleAnySlice(pos.Object)
		}

		var inp *vote.CreateI
		{
			inp = r.randomVote(fak, pos.Object[0])
		}

		var ctx context.Context
		{
			ctx = newCtx(key, use.Object[0].Public.Name)
		}

		var out *vote.CreateO
		{
			out, err = cli.Vote.Create(ctx, inp)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		for _, x := range out.Object {
			ids = append(ids, x.Intern.Id)
		}
	}

	var inp *vote.SearchI
	{
		inp = &vote.SearchI{}
	}

	for _, x := range ids {
		inp.Object = append(inp.Object, &vote.SearchI_Object{
			Intern: &vote.SearchI_Object_Intern{
				Id: x,
			},
		})
	}

	var out *vote.SearchO
	{
		out, err = cli.Vote.Search(context.Background(), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return out, nil
}

func (r *run) randomVote(fak *gofakeit.Faker, pos *post.SearchO_Object) *vote.CreateI {
	var opt []string
	{
		opt = []string{
			"true",
			"false",
		}
	}

	{
		fak.ShuffleAnySlice(opt)
	}

	var obj *vote.CreateI
	{
		obj = &vote.CreateI{
			Object: []*vote.CreateI_Object{
				{
					Public: &vote.CreateI_Object_Public{
						Claim:  pos.Intern.Id,
						Kind:   "stake",
						Option: opt[0],
						Value:  limStr(converter.FloatToString(fak.Float64Range(0.0001, 2.5)), 6),
					},
				},
			},
		}
	}

	return obj
}

func limStr(str string, lim int) string {
	if len(str) > lim {
		return str[:lim]
	}

	return str
}
