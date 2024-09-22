package fakeit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/xh3b4sd/tracer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (r *run) createPropose(key jwk.Key, use *user.SearchO) (*post.SearchO, error) {
	var err error

	var ids []string
	for i := 0; i < 20; i++ {
		{
			r.fak.ShuffleAnySlice(use.Object)
		}

		var inp *post.CreateI
		{
			inp = r.randomClaim()
		}

		var ctx context.Context
		{
			ctx = newCtx(key, use.Object[0].Public.Name)
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

	return out, nil
}

func (r *run) randomClaim() *post.CreateI {
	var con string
	{
		con = r.fak.HexUint(160)
	}

	var hsh string
	if r.fak.Float64() > 0.2 {
		hsh = r.fak.HexUint(256)
	}

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
		r.fak.ShuffleAnySlice(lab)
	}

	var tit string
	{
		tit = r.fak.BookTitle()
	}

	var par string
	{
		par = r.fak.Paragraph(r.fak.Number(1, 3), r.fak.Number(2, 5), r.fak.Number(10, 40), "\n\n")
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
						Chain:     "421614",
						Contract:  con,
						Expiry:    converter.TimeToString(time.Now().UTC().AddDate(0, r.fak.Number(1, 9), r.fak.Number(10, 30))),
						Hash:      hsh,
						Kind:      "claim",
						Labels:    strings.Join(lab[:r.fak.Number(1, 4)], ","),
						Lifecycle: string(objectlabel.LifecyclePropose),
						Meta:      "9,0",
						Text:      fmt.Sprintf("# %s\n\n%s\n\n%s", tit, par, strings.Join(lis[:r.fak.Number(2, 5)], "\n")),
						Token:     r.fak.RandomString([]string{"USDC", "UVX", "WETH"}),
					},
				},
			},
		}
	}

	return obj
}

func (r *run) ranLin(str string) string {
	var spl []string
	{
		spl = strings.Split(str, " ")
	}

	var ind int
	{
		ind = r.fak.Number(0, len(spl)-1)
	}

	var txt string
	{
		txt = r.fak.PetName()
	}

	if ind == 0 {
		txt = cases.Upper(language.English).String(txt)
	}

	var end string
	if ind == len(spl)-1 {
		end = "."
	}

	{
		spl[ind] = fmt.Sprintf("[%s](%s)", txt, r.fak.URL())
	}

	return strings.Join(spl, " ") + end
}
