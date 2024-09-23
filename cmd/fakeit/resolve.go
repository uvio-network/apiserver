package fakeit

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (r *run) createResolve(wal *wallet.SearchO, cla *post.SearchO) (*post.SearchO, error) {
	var err error

	var lis *post.SearchO
	{
		lis = witLif(cla, objectlabel.LifecyclePropose)
	}

	{
		r.fak.ShuffleAnySlice(lis)
		r.fak.ShuffleAnySlice(wal.Object)
	}

	all := map[string]*post.SearchO_Object{}
	for _, x := range lis.Object {
		all[x.Intern.Id] = x
	}

	var ids []string
	for i := 0; i < len(cla.Object)/2; i++ {
		var pro *post.SearchO_Object
		for _, v := range all {
			pro = v
			break
		}

		{
			delete(all, pro.Intern.Id)
		}

		var pos []*poststorage.Object
		{
			pos, err = r.dae.Sto().Post().SearchPost([]objectid.ID{objectid.ID(pro.Intern.Id)})
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		var exp time.Time
		{
			exp = pos[0].Expiry.AddDate(0, 0, r.fak.Number(3, 30))
		}

		var res *poststorage.Object
		{
			res, err = r.dae.Rec().Post().CreateResolve(pos[0], exp)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		{
			err = r.dae.Sto().Post().DeleteExpiry(pos)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		{
			pos[0].Expiry = time.Now().UTC()
		}

		// We are explicitely not adding the artificially expired propose to the
		// expiry queue, since the local fake setup is not fully implemented
		// onchain. If we were to expire the proposes here properly, the worker
		// handlers would not find proposes in the smart contracts of the local
		// blockchain node (anvil/hardhat).
		//
		// {
		// 	err = r.dae.Sto().Post().InternCreateExpiry(pos)
		// 	if err != nil {
		// 		tracer.Panic(tracer.Mask(err))
		// 	}
		// }

		{
			err = r.dae.Sto().Post().UpdatePost(pos)
			if err != nil {
				tracer.Panic(tracer.Mask(err))
			}
		}

		{
			ids = append(ids, string(res.ID))
		}

		if r.fak.Float64() > 0.2 {
			var hsh common.Hash
			{
				hsh = common.HexToHash(r.fak.HexUint(256))
			}

			var num []int
			{
				num = []int{
					1, 2, 4, 6, 8, 10,
				}
			}

			var add []common.Address
			for _, x := range wal.Object[:num[r.fak.Number(0, 5)]] {
				add = append(add, common.HexToAddress(x.Public.Address))
			}

			{
				err = r.dae.Rec().Post().UpdateResolve(res, hsh, add)
				if err != nil {
					tracer.Panic(tracer.Mask(err))
				}
			}
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

func witLif(pos *post.SearchO, lif objectlabel.DesiredLifecycle) *post.SearchO {
	var lis *post.SearchO
	{
		lis = &post.SearchO{}
	}

	for _, x := range pos.Object {
		if x.Public.Lifecycle == string(lif)+":"+string(objectlabel.LifecycleOnchain) {
			lis.Object = append(lis.Object, x)
		}
	}

	return lis
}
