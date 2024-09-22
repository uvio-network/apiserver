package fakeit

import (
	"context"
	"strings"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createVotePropose(key jwk.Key, use *user.SearchO, cla *post.SearchO) (*vote.SearchO, error) {
	var err error

	var ids []string
	for i := 0; i < 100; i++ {
		{
			r.fak.ShuffleAnySlice(use.Object)
			r.fak.ShuffleAnySlice(cla.Object)
		}

		var nam string
		{
			nam = use.Object[0].Public.Name
		}

		// As long as the claim is still "pending", only the owner can vote on it.
		if strings.HasSuffix(cla.Object[0].Public.Lifecycle, "pending") {
			nam = useUID(use, cla.Object[0].Intern.Owner).Public.Name
		}

		var inp *vote.CreateI
		{
			inp = r.randomVote(cla.Object[0])
		}

		var ctx context.Context
		{
			ctx = newCtx(key, nam)
		}

		var out *vote.CreateO
		{
			out, err = r.cli.Vote.Create(ctx, inp)
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
		out, err = r.cli.Vote.Search(context.Background(), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return out, nil
}

func (r *run) createVoteResolve(key jwk.Key, use *user.SearchO, cla *post.SearchO) (*vote.SearchO, error) {
	var err error

	{
		r.fak.ShuffleAnySlice(use.Object)
		r.fak.ShuffleAnySlice(cla.Object)
	}

	var ids []string
	for i := 0; i < len(cla.Object); i++ {
		if strings.HasSuffix(cla.Object[i].Public.Lifecycle, "pending") {
			continue
		}

		var num int
		{
			num = r.fak.Number(0, len(cla.Object[i].Extern.Samples))
		}

		if num == 0 {
			continue
		}

		var nam []string
		for _, v := range cla.Object[i].Extern.Samples {
			if len(nam) < num {
				nam = append(nam, useUID(use, v).Public.Name)
			}
		}

		for _, x := range nam {
			var inp *vote.CreateI
			{
				inp = r.randomVote(cla.Object[i])
			}

			var ctx context.Context
			{
				ctx = newCtx(key, x)
			}

			var out *vote.CreateO
			{
				out, err = r.cli.Vote.Create(ctx, inp)
				if err != nil {
					tracer.Panic(tracer.Mask(err))
				}
			}

			for _, x := range out.Object {
				ids = append(ids, x.Intern.Id)
			}
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
		out, err = r.cli.Vote.Search(context.Background(), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return out, nil
}

func (r *run) randomVote(cla *post.SearchO_Object) *vote.CreateI {
	var pro bool
	if strings.HasPrefix(cla.Public.Lifecycle, "propose") {
		pro = true
	}

	var res bool
	if strings.HasPrefix(cla.Public.Lifecycle, "resolve") {
		res = true
	}

	var hsh string
	if r.fak.Float64() > 0.2 {
		hsh = r.fak.HexUint(256)
	}

	var kin string
	if pro {
		kin = "stake"
	} else if res {
		kin = "truth"
	}

	var opt []string
	{
		opt = []string{
			"true",
			"false",
		}
	}

	var val string
	if pro {
		val = limStr(converter.FloatToString(r.fak.Float64Range(0.0001, 2.5)), 6)
	} else if res {
		val = "1"
	}

	var obj *vote.CreateI
	{
		obj = &vote.CreateI{
			Object: []*vote.CreateI_Object{
				{
					Public: &vote.CreateI_Object_Public{
						Claim:     cla.Intern.Id,
						Hash:      hsh,
						Kind:      kin,
						Lifecycle: string(objectlabel.LifecycleOnchain),
						Option:    r.fak.RandomString(opt),
						Value:     val,
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
