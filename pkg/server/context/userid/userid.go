package userid

import (
	"context"

	"github.com/xh3b4sd/objectid"
)

type t string

var k t = "userid"

func NewContext(ctx context.Context, v objectid.ID) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context) objectid.ID {
	v, _ := ctx.Value(k).(objectid.ID)
	return v
}
