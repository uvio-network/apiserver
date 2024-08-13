package poststorage

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/labelname"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Chain     string      `json:"chain"`
	Created   time.Time   `json:"created"`
	Expiry    time.Time   `json:"expiry"`
	ID        objectid.ID `json:"id"`
	Kind      string      `json:"kind"`
	Labels    []string    `json:"labels"`
	Lifecycle string      `json:"lifecycle"`
	Meta      string      `json:"meta"`
	Owner     objectid.ID `json:"owner"`
	Parent    objectid.ID `json:"parent"`
	Text      string      `json:"text"`
	Token     string      `json:"token"`
	Tree      objectid.ID `json:"tree"`
	Votes     []float64   `json:"votes"`
}

func (o *Object) Verify() error {
	{
		if o.Kind == "claim" && o.Chain == "" {
			return tracer.Mask(ClaimChainEmptyError)
		}
		if o.Kind == "comment" && o.Chain != "" {
			return tracer.Mask(ClaimChainInvalidError)
		}
	}

	{
		if o.Kind == "claim" && o.Expiry.IsZero() {
			return tracer.Mask(ClaimExpiryEmptyError)
		}

		// When creating posts of kind "claim", we want to ensure that claims cannot
		// be created with expiries that point to the past. During claim creation,
		// the post IDs are only allocated once the input data got verified. During
		// claim updates, the post expiry may very well be in the past at the time
		// of the update being processed. During those update processes we do not
		// want to run the check below again, assuming nobody from the outside could
		// change the initial expiry anymore.
		if o.ID == "" && o.Kind == "claim" && o.Expiry.Compare(time.Now().UTC()) != +1 {
			return tracer.Mask(ClaimExpiryPastError)
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
		if o.Kind == "claim" && o.Lifecycle != "adjourn" && o.Lifecycle != "dispute" && o.Lifecycle != "nullify" && o.Lifecycle != "propose" && o.Lifecycle != "resolve" {
			return tracer.Maskf(ClaimLifecycleInvalidError, o.Lifecycle)
		}
		if o.Kind == "comment" && o.Lifecycle != "" {
			return tracer.Maskf(CommentLifecycleInvalidError, o.Lifecycle)
		}
	}

	{
		if o.Kind == "claim" && o.Meta == "" {
			return tracer.Mask(ClaimMetaEmptyError)
		}
		if o.Kind == "comment" && o.Meta != "" {
			return tracer.Mask(ClaimMetaInvalidError)
		}
	}

	{
		// Any claim with lifecycle other than "propose" must reference a parent.
		if o.Kind == "claim" && o.Lifecycle != "propose" && o.Parent == "" {
			return tracer.Maskf(ClaimParentEmptyError, o.Lifecycle)
		}
		// Any claim with lifecycle "propose" must not reference a parent.
		if o.Kind == "claim" && o.Lifecycle == "propose" && o.Parent != "" {
			return tracer.Maskf(ClaimParentInvalidError, o.Lifecycle)
		}
		// Any comment must reference its parent claim.
		if o.Kind == "comment" && o.Parent == "" {
			return tracer.Maskf(CommentParentEmptyError, o.Lifecycle)
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
		if o.Kind == "claim" && (o.Lifecycle == "adjourn" || o.Lifecycle == "dispute" || o.Lifecycle == "nullify" || o.Lifecycle == "propose") && tok == "" {
			return tracer.Mask(PostTokenEmptyError)
		}
		// Any claim on which you cannot stake reputation must not specify a staking
		// token.
		if o.Kind == "claim" && o.Lifecycle == "resolve" && tok != "" {
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
