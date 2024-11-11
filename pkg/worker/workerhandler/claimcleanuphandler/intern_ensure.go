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
	oneDay  = 24 * time.Hour
	oneWeek = 24 * time.Hour * 7
)

func (h *InternHandler) Ensure(tas *task.Task) error {
	var err error

	var now time.Time
	{
		now = time.Now().UTC()
	}

	var day time.Time
	var wee time.Time
	{
		day = now.Add(-oneDay)
		wee = now.Add(-oneWeek)
	}

	// Look for all claims of lifecycle phase created using unix seconds. We want
	// to ignore all claims that are pending within their first 24 hours of
	// existance. After that we assumme that those claims should be cleaned up for
	// good.
	//
	//             |       cleanup      | ignore |
	//     --------x--------------------x--------x----
	//          -1 week              -1 day     now
	//
	var pos poststorage.Slicer
	{
		pos, err = h.sto.Post().SearchTime(objectlabel.LifecycleCreated, wee.Unix(), day.Unix())
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
