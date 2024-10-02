package votereconciler

import (
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) UpdateHash(vot []*votestorage.Object, hsh []string) ([]*votestorage.Object, error) {
	for i := range vot {
		if !vot[i].Lifecycle.Pending() {
			return nil, tracer.Maskf(VoteUpdateHashError, "%s = %s", vot[i].ID, vot[i].Lifecycle.Hash)
		}

		{
			vot[i].Lifecycle.Hash = append(vot[i].Lifecycle.Hash, hsh[i])
		}
	}

	return vot, nil
}
