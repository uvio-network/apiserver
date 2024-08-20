package votereconciler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateHash(vot []*votestorage.Object, has []string) ([]*votestorage.Object, error) {
	var err error

	for i := range vot {
		if vot[i].Lifecycle.Hash != "" {
			return nil, tracer.Maskf(VoteUpdateHashError, "%s=%s", vot[i].ID, vot[i].Lifecycle.Hash)
		}

		{
			vot[i].Lifecycle.Hash = has[i]
		}

		if vot[i].Kind == "stake" {
			var cla *poststorage.Object
			{
				cla, err = r.searchClaim(vot[i].Claim)
				if err != nil {
					return nil, tracer.Mask(err)
				}
			}

			var now time.Time
			{
				now = time.Now().UTC()
			}

			err = r.updateStaked(vot[i], cla.Token, now)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}
	}

	return vot, nil
}

func (r *Redis) updateStaked(vot *votestorage.Object, tok string, now time.Time) error {
	var err error

	var use []*userstorage.Object
	{
		use, err = r.sto.User().SearchUser([]objectid.ID{vot.Owner})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for j := range use {
		use[j].Staked.Data[tok] += vot.Value
		use[j].Staked.Time = now
	}

	{
		err = r.sto.User().UpdateUser(use)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
