package votestorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/format/hexencoding"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Claim     objectid.ID           `json:"claim,omitempty"`
	Created   time.Time             `json:"created,omitempty"`
	ID        objectid.ID           `json:"id,omitempty"`
	Kind      string                `json:"kind,omitempty"`
	Lifecycle objectfield.Lifecycle `json:"lifecycle,omitempty"`
	Meta      string                `json:"meta,omitempty"`
	Option    bool                  `json:"option"` // no omitempty because we want to see "false" in the JSON
	Owner     objectid.ID           `json:"owner,omitempty"`
	Value     float64               `json:"value,omitempty"`
}

func (o *Object) Verify() error {
	{
		if o.Kind != "stake" && o.Kind != "truth" {
			return tracer.Maskf(VoteKindInvalidError, o.Kind)
		}
	}

	{
		if !o.Lifecycle.Is(objectlabel.LifecycleOnchain) {
			return tracer.Maskf(VoteLifecycleInvalidError, string(o.Lifecycle.Data))
		}
		for _, x := range o.Lifecycle.Hash {
			if x != "" && !hexencoding.Verify(x) {
				return tracer.Maskf(VoteHashFormatError, x)
			}
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
