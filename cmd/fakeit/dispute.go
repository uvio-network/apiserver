package fakeit

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/summary"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createDispute(key jwk.Key, use *user.SearchO, cla *post.SearchO) (*post.SearchO, error) {
	var err error

	var lis *post.SearchO
	{
		lis = witLif(cla, objectlabel.LifecycleResolve)
	}

	{
		r.fak.ShuffleAnySlice(lis)
		r.fak.ShuffleAnySlice(use.Object)
	}

	all := map[string]*post.SearchO_Object{}
	for _, x := range lis.Object {
		all[x.Intern.Id] = x
	}

	var ids []string
	for i := 0; i < len(lis.Object); i++ {
		var res *post.SearchO_Object

		for _, v := range all {
			res = v
			break
		}

		{
			delete(all, res.Intern.Id)
		}

		var oid []objectid.ID
		{
			oid = []objectid.ID{
				objectid.ID(res.Public.Parent),
				objectid.ID(res.Intern.Id),
			}
		}

		var pos []*poststorage.Object
		{
			pos, err = r.dae.Sto().Post().SearchPost(oid)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		// The second post object here is the resolve. We can only go ahead and
		// create a dispute for it if this resolve has a valid resolution.
		if !summary.Verify(pos[1]) {
			continue
		}

		var inp *post.CreateI
		{
			inp = r.randomDispute(pos[0], res)
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

	return witLif(out, objectlabel.LifecycleDispute), nil
}

func (r *run) randomDispute(pro *poststorage.Object, res *post.SearchO_Object) *post.CreateI {
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
		tit = r.fak.BeerName()
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
						Chain:     pro.Chain,
						Contract:  pro.Contract,
						Expiry:    converter.TimeToString(time.Now().UTC().AddDate(0, r.fak.Number(1, 9), r.fak.Number(10, 30))),
						Hash:      hsh,
						Kind:      "claim",
						Labels:    strings.Join(lab[:r.fak.Number(1, 4)], ","),
						Lifecycle: string(objectlabel.LifecycleDispute),
						Parent:    res.Intern.Id,
						Text:      fmt.Sprintf("# %s\n\n%s\n\n%s", tit, par, strings.Join(lis[:r.fak.Number(2, 5)], "\n")),
						Token:     pro.Token,
					},
				},
			},
		}
	}

	return obj
}
