package votestorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/format/hexencoding"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Chain     string      `json:"chain"`
	Claim     objectid.ID `json:"claim"`
	Created   time.Time   `json:"created"`
	Hash      string      `json:"hash"`
	ID        objectid.ID `json:"id"`
	Kind      string      `json:"kind"`
	Lifecycle string      `json:"lifecycle"`
	Meta      string      `json:"meta"`
	Option    bool        `json:"option"`
	Owner     objectid.ID `json:"owner"`
	Value     float64     `json:"value"`
}

func (o *Object) Verify() error {
	{
		if o.Chain == "" {
			return tracer.Mask(VoteChainEmptyError)
		}
	}

	{
		if !hexencoding.Verify(o.Hash) {
			return tracer.Mask(VoteHashFormatError)
		}
	}

	{
		if o.Kind != "stake" && o.Kind != "truth" {
			return tracer.Maskf(VoteKindInvalidError, o.Kind)
		}
	}

	{
		// Any vote without hash must be in lifecycle phase "pending".
		if o.Hash == "" && o.Lifecycle != "pending" {
			return tracer.Maskf(VoteHashPendingError, o.Lifecycle)
		}
		// Any vote with hash must be in lifecycle phase "onchain".
		if o.Hash != "" && o.Lifecycle != "onchain" {
			return tracer.Maskf(VoteHashLifecycleError, o.Lifecycle)
		}
		if o.Lifecycle != "onchain" && o.Lifecycle != "pending" {
			return tracer.Maskf(VoteLifecycleInvalidError, o.Lifecycle)
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
