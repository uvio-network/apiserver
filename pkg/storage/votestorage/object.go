package votestorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/format/hexencoding"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Chain     string                `json:"chain"`
	Claim     objectid.ID           `json:"claim"`
	Created   time.Time             `json:"created"`
	ID        objectid.ID           `json:"id"`
	Kind      string                `json:"kind"`
	Lifecycle objectfield.Lifecycle `json:"lifecycle"`
	Meta      string                `json:"meta"`
	Option    bool                  `json:"option"`
	Owner     objectid.ID           `json:"owner"`
	Value     float64               `json:"value"`
}

func (o *Object) Verify() error {
	{
		if o.Chain == "" {
			return tracer.Mask(VoteChainEmptyError)
		}
	}

	{
		if o.Kind != "stake" && o.Kind != "truth" {
			return tracer.Maskf(VoteKindInvalidError, o.Kind)
		}
	}

	{
		if !o.Lifecycle.Is(objectlabel.LifecycleOnchain) {
			return tracer.Maskf(VoteLifecycleInvalidError, o.Lifecycle.String())
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
