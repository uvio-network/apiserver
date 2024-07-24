package poststorage

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Created   time.Time   `json:"created"`
	Expiry    time.Time   `json:"expiry"`
	ID        objectid.ID `json:"id"`
	Kind      string      `json:"kind"`
	Lifecycle string      `json:"lifecycle"`
	Option    bool        `json:"option"`
	Owner     objectid.ID `json:"owner"`  // TODO verify post creator once auth is setup
	Parent    objectid.ID `json:"parent"` // TODO must exist if given
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
	}

	{
		tok := strings.TrimSpace(o.Token)

		if tok == "" {
			return tracer.Mask(PostTextEmptyError)
		}
	}

	// {
	// 	if o.User == "" {
	// 		return tracer.Mask(runtime.UserIDEmptyError)
	// 	}
	// }

	return nil
}
