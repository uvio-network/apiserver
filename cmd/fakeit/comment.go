package fakeit

import (
	"context"
	"fmt"
	"strings"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createComment(key jwk.Key, use *user.SearchO, cla *post.SearchO, vot *vote.SearchO) (*post.SearchO, error) {
	var err error

	var ids []string
	for i := 0; i < 15; i++ {
		{
			r.fak.ShuffleAnySlice(vot.Object)
		}

		var inp *post.CreateI
		{
			inp = r.randomComment(claVot(cla, vot.Object[0].Public.Claim))
		}

		if inp == nil {
			continue
		}

		var ctx context.Context
		{
			ctx = newCtx(key, useUID(use, vot.Object[0].Intern.Owner).Public.Name)
		}

		var out *post.CreateO
		{
			out, err = r.cli.Post.Create(ctx, inp)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		for _, x := range out.Object {
			ids = append(ids, x.Intern.Id)
		}
	}

	var inp *post.SearchI
	{
		inp = &post.SearchI{}
	}

	for _, x := range ids {
		inp.Object = append(inp.Object, &post.SearchI_Object{
			Intern: &post.SearchI_Object_Intern{
				Id: x,
			},
		})
	}

	var out *post.SearchO
	{
		out, err = r.cli.Post.Search(context.Background(), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	{
		out = remPar(out)
	}

	return out, nil
}

func (r *run) randomComment(cla *post.SearchO_Object) *post.CreateI {
	if cla == nil {
		return nil
	}

	var tit string
	{
		tit = r.fak.BookTitle()
	}

	var par string
	{
		par = r.fak.Paragraph(r.fak.Number(1, 3), r.fak.Number(1, 5), r.fak.Number(5, 40), "\n\n")
	}

	var lis []string
	{
		lis = []string{
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
			"* " + r.ranLin(r.fak.Sentence(r.fak.Number(3, 9))),
		}
	}

	{
		r.fak.ShuffleAnySlice(lis)
	}

	var obj *post.CreateI
	{
		obj = &post.CreateI{
			Object: []*post.CreateI_Object{
				{
					Public: &post.CreateI_Object_Public{
						Kind:   "comment",
						Parent: cla.Intern.Id,
						Text:   fmt.Sprintf("# %s\n\n%s\n\n%s", tit, par, strings.Join(lis[:r.fak.Number(1, 5)], "\n")),
					},
				},
			},
		}
	}

	return obj
}

func claVot(cla *post.SearchO, cid string) *post.SearchO_Object {
	for _, x := range cla.Object {
		if strings.Contains(x.Public.Lifecycle, "pending") {
			continue
		}

		if x.Intern.Id == cid {
			return x
		}
	}

	return nil
}

func remPar(pos *post.SearchO) *post.SearchO {
	var lis *post.SearchO
	{
		lis = &post.SearchO{}
	}

	for _, x := range pos.Object {
		// Only keep the comments and ignore the claims.
		if x.Public.Kind != "claim" {
			lis.Object = append(lis.Object, x)
		}
	}

	return lis
}
