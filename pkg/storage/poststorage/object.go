package poststorage

import (
	"strings"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Created time.Time   `json:"created"`
	Expiry  time.Time   `json:"expiry"` // TODO expiry must be greater than created
	ID      objectid.ID `json:"post"`
	Kind    string      `json:"kind"` // TODO kind is either claim or comment
	Option  bool        `json:"option"`
	Owner   objectid.ID `json:"owner"` // TODO verify post creator once auth is setup
	Stake   float64     `json:"stake"` // TODO stake must not be empty
	Text    string      `json:"text"`
	Token   string      `json:"token"`
}

func (o *Object) Verify() error {
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
