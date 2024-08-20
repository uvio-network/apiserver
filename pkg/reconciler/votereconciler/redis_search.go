package votereconciler

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/xh3b4sd/tracer"
)

func (r *Redis) searchClaim(cla objectid.ID) (*poststorage.Object, error) {
	var err error

	var obj []*poststorage.Object
	{
		obj, err = r.sto.Post().SearchPost([]objectid.ID{cla})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// Every vote object must reference an existing claim object.
	if len(obj) != 1 {
		return nil, tracer.Maskf(VoteClaimNotFoundError, cla.String())
	}

	return obj[0], nil
}
