package votereconciler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) UpdateHash(use objectid.ID, ids []objectid.ID, has []string) ([]*votestorage.Object, error) {
	var err error
	var pos []*votestorage.Object

	if len(ids) != len(has) {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "%d != %d", len(ids), len(has))
	}

	{
		pos, err = r.sto.Vote().SearchVote(ids)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	for i := range pos {
		if pos[i].Lifecycle.Hash != "" {
			return nil, tracer.Maskf(VoteUpdateHashError, "%s=%s", pos[i].ID, pos[i].Lifecycle.Hash)
		}

		if pos[i].Owner != use {
			return nil, tracer.Maskf(runtime.UserNotOwnerError, "%s=%s", pos[i].Owner, use)
		}

		{
			pos[i].Lifecycle.Hash = has[i]
		}
	}

	return pos, nil
}
