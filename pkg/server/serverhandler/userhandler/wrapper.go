package userhandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/user"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *user.CreateI) (*user.CreateO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		if len(req.Object) > 100 {
			return nil, tracer.Mask(runtime.QueryObjectLimitError)
		}

		for _, x := range req.Object {
			if !x.ProtoReflect().IsValid() {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	{
		for _, x := range req.Object {
			if x.Intern != nil {
				return nil, tracer.Maskf(runtime.QueryObjectInvalidError, "intern must be empty")
			}
			if !x.Public.ProtoReflect().IsValid() {
				return nil, tracer.Maskf(runtime.QueryObjectInvalidError, "public must not be empty")
			}
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *user.DeleteI) (*user.DeleteO, error) {
	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *user.SearchI) (*user.SearchO, error) {
	{
		if len(req.Object) == 0 {
			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
		}

		if len(req.Object) > 100 {
			return nil, tracer.Mask(runtime.QueryObjectLimitError)
		}

		for _, x := range req.Object {
			if !x.ProtoReflect().IsValid() {
				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
			}
		}
	}

	{
		for _, x := range req.Object {
			i := x.Intern.ProtoReflect().IsValid()
			p := x.Public.ProtoReflect().IsValid()

			if i && !p || !i && p {
				return nil, tracer.Maskf(runtime.QueryObjectConflictError, "intern and public must not be used together")
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *user.UpdateI) (*user.UpdateO, error) {
	return w.han.Update(ctx, req)
}
