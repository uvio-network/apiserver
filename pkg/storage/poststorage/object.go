package poststorage

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/hexencoding"
	"github.com/uvio-network/apiserver/pkg/format/labelname"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Chain     string                `json:"chain,omitempty"`
	Contract  string                `json:"contract,omitempty"`
	Created   time.Time             `json:"created,omitempty"`
	Expiry    time.Time             `json:"expiry,omitempty"`
	ID        objectid.ID           `json:"id,omitempty"`
	Kind      string                `json:"kind,omitempty"`
	Labels    []string              `json:"labels,omitempty"`
	Lifecycle objectfield.Lifecycle `json:"lifecycle,omitempty"`
	Meta      string                `json:"meta,omitempty"`
	Owner     objectid.ID           `json:"owner,omitempty"`
	Parent    objectid.ID           `json:"parent,omitempty"`
	Samples   map[string]string     `json:"samples,omitempty"`
	Text      string                `json:"text,omitempty"`
	Token     string                `json:"token,omitempty"`
	Tree      objectid.ID           `json:"tree,omitempty"`
	Votes     []float64             `json:"votes,omitempty"` // TODO rename to Summary
}

func (o *Object) Verify() error {
	{
		if o.Kind == "claim" && o.Chain == "" {
			return tracer.Mask(ClaimChainEmptyError)
		}
		if o.Kind == "comment" && o.Chain != "" {
			return tracer.Maskf(ClaimChainInvalidError, o.Chain)
		}
	}

	{
		if o.Kind == "claim" && o.Contract == "" {
			return tracer.Mask(ClaimContractEmptyError)
		}
	}

	{
		if o.Kind == "claim" && !o.Lifecycle.Is(objectlabel.LifecycleBalance) {
			// All claim objects must have an expiry, except the internal claim type
			// "balance". Balance claims are only for us internally to ensure all user
			// balances get updated once the original propose is settled onchain.
			if o.Expiry.IsZero() {
				return tracer.Mask(ClaimExpiryEmptyError)
			}

			// When creating posts of kind "claim", we want to ensure that claims cannot
			// be created with expiries that point to the past. During claim creation,
			// the post IDs are only allocated once the input data got verified. During
			// claim updates, the post expiry may very well be in the past at the time
			// of the update being processed. During those update processes we do not
			// want to run the check below again, assuming nobody from the outside could
			// change the initial expiry anymore.
			if o.ID == "" && o.Expiry.Compare(time.Now().UTC()) != +1 {
				return tracer.Maskf(ClaimExpiryPastError, "%d", o.Expiry.Unix())
			}
		}
	}

	{
		if o.Kind != "claim" && o.Kind != "comment" {
			return tracer.Maskf(PostKindInvalidError, o.Kind)
		}
	}

	{
		if o.Kind == "comment" && len(o.Labels) != 0 {
			return tracer.Mask(CommentLabelsInvalidError)
		}
		if o.Kind == "claim" && len(o.Labels) == 0 {
			return tracer.Mask(ClaimLabelsEmptyError)
		}
		if o.Kind == "claim" && len(o.Labels) > 4 {
			return tracer.Mask(ClaimLabelsLimitError)
		}
		if o.Kind == "claim" && generic.Duplicate(o.Labels) {
			return tracer.Mask(ClaimLabelsUniqueError)
		}
		if o.Kind == "claim" && !labelname.Verify(o.Labels) {
			return tracer.Maskf(ClaimLabelsFormatError, "%v", o.Labels)
		}
	}

	{
		if o.Kind == "claim" && !o.Lifecycle.Is(objectlabel.LifecycleBalance, objectlabel.LifecycleDispute, objectlabel.LifecyclePropose, objectlabel.LifecycleResolve) {
			return tracer.Maskf(ClaimLifecycleInvalidError, string(o.Lifecycle.Data))
		}
		if o.Kind == "comment" && !o.Lifecycle.Empty() {
			return tracer.Maskf(CommentLifecycleInvalidError, string(o.Lifecycle.Data))
		}
		for _, x := range o.Lifecycle.Hash {
			if x != "" && !hexencoding.Verify(x) {
				return tracer.Maskf(ClaimHashFormatError, x)
			}
		}
	}

	{
		// Any claim with lifecycle other than "propose" must reference a parent.
		if o.Kind == "claim" && !o.Lifecycle.Is(objectlabel.LifecyclePropose) && o.Parent == "" {
			return tracer.Maskf(ClaimParentEmptyError, string(o.Lifecycle.Data))
		}
		// Any claim with lifecycle "propose" must not reference a parent.
		if o.Kind == "claim" && o.Lifecycle.Is(objectlabel.LifecyclePropose) && o.Parent != "" {
			return tracer.Maskf(ClaimParentInvalidError, string(o.Lifecycle.Data))
		}
		// Any comment must reference its parent claim.
		if o.Kind == "comment" && o.Parent == "" {
			return tracer.Maskf(CommentParentEmptyError, string(o.Lifecycle.Data))
		}
	}

	{
		if o.Owner == "" {
			return tracer.Mask(PostOwnerEmptyError)
		}
	}

	{
		txt := strings.TrimSpace(o.Text)

		if txt == "" {
			return tracer.Mask(PostTextEmptyError)
		}
		if len(txt) < 100 {
			return tracer.Maskf(PostTextLengthError, "%d", len(txt))
		}
		if len(txt) > 5000 {
			return tracer.Maskf(PostTextLengthError, "%d", len(txt))
		}
	}

	{
		tok := strings.TrimSpace(o.Token)

		// Any claim on which you can stake reputation must specify a staking token.
		if o.Kind == "claim" && o.Lifecycle.Is(objectlabel.LifecycleDispute, objectlabel.LifecyclePropose) && tok == "" {
			return tracer.Mask(PostTokenEmptyError)
		}
		// Any claim on which you cannot stake reputation must not specify a staking
		// token.
		if o.Kind == "claim" && o.Lifecycle.Is(objectlabel.LifecycleResolve) && tok != "" {
			return tracer.Mask(PostTokenInvalidError)
		}
		// Any comment must not specify a staking token.
		if o.Kind == "comment" && tok != "" {
			return tracer.Mask(PostTokenInvalidError)
		}
	}

	{
		if o.Kind == "claim" && len(o.Votes) != 4 {
			return tracer.Maskf(runtime.ExecutionFailedError, "vote summary for claims must contain 4 parts")
		}

		if o.Kind == "comment" && len(o.Votes) != 2 {
			return tracer.Maskf(runtime.ExecutionFailedError, "vote summary for comments must contain 2 parts")
		}
	}

	return nil
}
