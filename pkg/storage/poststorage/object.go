package poststorage

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/format/labelname"
	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Created   time.Time   `json:"created"`
	Expiry    time.Time   `json:"expiry"`
	ID        objectid.ID `json:"id"`
	Kind      string      `json:"kind"`
	Labels    []string    `json:"labels"`
	Lifecycle string      `json:"lifecycle"`
	Option    bool        `json:"option"`
	Owner     objectid.ID `json:"owner"`
	Parent    objectid.ID `json:"parent"`
	Stake     float64     `json:"stake"`
	Text      string      `json:"text"`
	Token     string      `json:"token"`
	Tree      objectid.ID `json:"tree"`
}

func (o *Object) Verify() error {
	{
		if o.Expiry.IsZero() {
			return tracer.Mask(PostExpiryEmptyError)
		}

		// When creating posts of kind "claim", we want to ensure that claims cannot
		// be created with expiries that point to the past. During claim creation,
		// the post IDs are only allocated once the input data got verified. During
		// claim updates, the post expiry may very well be in the past at the time
		// of the update being processed. During those update processes we do not
		// want to run the check below again, assuming nobody from the outside could
		// change the initial expiry anymore.
		if o.ID == "" && o.Kind == "claim" && o.Expiry.Compare(time.Now().UTC()) != +1 {
			return tracer.Mask(PostExpiryPastError)
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
			return tracer.Mask(ClaimLabelsFormatError)
		}
	}

	{
		if o.Kind == "claim" && o.Lifecycle != "propose" && o.Lifecycle != "resolve" && o.Lifecycle != "dispute" && o.Lifecycle != "nullify" {
			return tracer.Maskf(ClaimLifecycleInvalidError, o.Lifecycle)
		}
		if o.Kind == "comment" && o.Lifecycle != "" {
			return tracer.Maskf(CommentLifecycleInvalidError, o.Lifecycle)
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
		if o.Stake <= 0 {
			return tracer.Mask(PostStakeInvalidError)
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
		if len(txt) > 2500 {
			return tracer.Maskf(PostTextLengthError, "%d", len(txt))
		}
	}

	{
		tok := strings.TrimSpace(o.Token)

		if tok == "" {
			return tracer.Mask(PostTextEmptyError)
		}
	}

	{
		if o.Owner == "" {
			return tracer.Mask(PostOwnerEmptyError)
		}
	}

	return nil
}
