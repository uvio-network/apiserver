package claimcleanuphandler

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

const (
	oneDay = 24 * time.Hour * 7 // TODO revert to one day once the worker ran on testnet
)

func (h *InternHandler) Ensure(tas *task.Task) error {
	var err error

	var now time.Time
	var day time.Time
	{
		now = time.Now().UTC()
		day = now.Add(-oneDay)
	}

	// Look for all claims of lifecycle phase propose using microsecond scores. We
	// have to use microseconds because the underlying claims are indexed using
	// their IDs as scores, where those IDs are in randomized microsecond format
	// based on claim creation time.
	var pos poststorage.Slicer
	{
		pos, err = h.sto.Post().SearchTime(objectlabel.LifecyclePropose, day.UnixMicro(), now.UnixMicro())
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		pos = pos.ObjectPending()
	}

	// It may very well happen that there are no pending claims to be deleted. In
	// that case we stop processing here.
	if len(pos) == 0 {
		return nil
	}

	for _, x := range pos {
		var vot votestorage.Slicer
		{
			vot, err = h.sto.Vote().SearchClaim([]objectid.ID{x.ID})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.sto.Vote().DeleteVote(vot)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var tre poststorage.Slicer
		{
			tre, err = h.sto.Post().SearchTree([]objectid.ID{x.Tree})
			if err != nil {
				return tracer.Mask(err)
			}
		}

		{
			err = h.sto.Post().DeletePost(tre)
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
