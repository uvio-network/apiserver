package votestorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Claim   objectid.ID `json:"claim"`
	Created time.Time   `json:"created"`
	ID      objectid.ID `json:"id"`
	Kind    string      `json:"kind"`
	Option  bool        `json:"option"`
	Owner   objectid.ID `json:"owner"`
	Value   float64     `json:"value"`
}

func (o *Object) Verify() error {
	{
		if o.Kind != "stake" && o.Kind != "truth" {
			return tracer.Maskf(VoteKindInvalidError, o.Kind)
		}
	}

	{
		if o.Owner == "" {
			return tracer.Mask(VoteOwnerEmptyError)
		}
	}

	{
		if o.Kind == "stake" && o.Value <= 0 {
			return tracer.Mask(StakeValueInvalidError)
		}

		if o.Kind == "truth" && o.Value != 1 {
			return tracer.Mask(TruthValueInvalidError)
		}
	}

	return nil
}
