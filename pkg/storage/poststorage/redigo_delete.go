package poststorage

import (
	"github.com/uvio-network/apiserver/pkg/format/storageformat"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) DeleteExpiry(inp []*Object) error {
	var err error

	for i := range inp {
		if inp[i].Kind == "claim" {
			err = r.red.Sorted().Delete().Value(posExp(inp[i].Lifecycle.Data), inp[i].ID.String())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}

func (r *Redigo) DeletePost(inp []*Object) error {
	var err error

	for i := range inp {
		if inp[i].Kind == "comment" {
			err = r.red.Sorted().Delete().Score(posOwnCom(inp[i].Owner, inp[i].Parent), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}

			err = r.red.Sorted().Delete().Score(posCom(inp[i].Parent), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		if inp[i].Kind == "claim" {
			err = r.red.Sorted().Delete().Score(posLif(inp[i].Lifecycle.Data), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}

			err = r.red.Sorted().Delete().Value(posExp(inp[i].Lifecycle.Data), inp[i].ID.String())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Sorted().Delete().Score(posOwn(inp[i].Owner), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Sorted().Delete().Score(posTre(inp[i].Tree), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		for _, x := range inp[i].Labels {
			err = r.red.Sorted().Delete().Score(posLab(x), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = r.red.Sorted().Delete().Value(storageformat.PostCreated, inp[i].ID.String())
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			_, err = r.red.Simple().Delete().Multi(posObj(inp[i].ID))
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
