package usercreated

import (
	"context"
)

type t string

var k t = "usercreated"

func NewContext(ctx context.Context, v bool) context.Context {
	return context.WithValue(ctx, k, v)
}

func FromContext(ctx context.Context) bool {
	v, _ := ctx.Value(k).(bool)
	return v
}
