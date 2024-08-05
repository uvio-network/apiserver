package fakeit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/xh3b4sd/tracer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (r *run) createPost(cli Client, key jwk.Key, fak *gofakeit.Faker, use *user.SearchO) (*post.SearchO, error) {
	var err error

	var ids []string
	for i := 0; i < 10; i++ {
		{
			fak.ShuffleAnySlice(use.Object)
		}

		var inp *post.CreateI
		{
			inp = r.randomPost(fak)
		}

		var ctx context.Context
		{
			ctx = newCtx(key, use.Object[0].Public.Name)
		}

		var out *post.CreateO
		{
			out, err = cli.Post.Create(ctx, inp)
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
		out, err = cli.Post.Search(context.Background(), inp)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return out, nil
}

func (r *run) randomPost(fak *gofakeit.Faker) *post.CreateI {
	var lab []string
	{
		lab = []string{
			"Finance",
			"Politics",
			"Crypto",
			"Economy",
			"Startups",
			"Sports",
			"Weather",
			"Science",
			"Technology",
		}
	}

	{
		fak.ShuffleAnySlice(lab)
	}

	var tit string
	{
		tit = fak.BookTitle()
	}

	var par string
	{
		par = fak.Paragraph(fak.Number(1, 3), fak.Number(2, 5), fak.Number(10, 50), "\n\n")
	}

	var lis []string
	{
		lis = []string{
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
			"* " + ranLin(fak, fak.Sentence(fak.Number(3, 9))),
		}
	}

	{
		fak.ShuffleAnySlice(lis)
	}

	var obj *post.CreateI
	{
		obj = &post.CreateI{
			Object: []*post.CreateI_Object{
				{
					Public: &post.CreateI_Object_Public{
						Expiry:    converter.TimeToString(time.Now().AddDate(0, fak.Number(1, 9), fak.Number(10, 30))),
						Kind:      "claim",
						Labels:    strings.Join(lab[:fak.Number(1, 4)], ","),
						Lifecycle: "propose",
						Text:      fmt.Sprintf("# %s\n\n%s\n\n%s", tit, par, strings.Join(lis[:fak.Number(2, 5)], "\n")),
						Token:     "USDC",
					},
				},
			},
		}
	}

	return obj
}

func ranLin(fak *gofakeit.Faker, str string) string {
	var spl []string
	{
		spl = strings.Split(str, " ")
	}

	var ind int
	{
		ind = fak.Number(0, len(spl)-1)
	}

	var txt string
	{
		txt = fak.PetName()
	}

	if ind == 0 {
		txt = cases.Upper(language.English).String(txt)
	}

	var end string
	if ind == len(spl)-1 {
		end = "."
	}

	{
		spl[ind] = fmt.Sprintf("[%s](%s)", txt, fak.URL())
	}

	return strings.Join(spl, " ") + end
}
