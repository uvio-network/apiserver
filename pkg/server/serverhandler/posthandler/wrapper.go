package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *post.CreateI) (*post.CreateO, error) {
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

	{
		if userid.FromContext(ctx) == "" {
			return nil, tracer.Mask(runtime.UserAuthError)
		}
	}

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *post.DeleteI) (*post.DeleteO, error) {
	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
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
			s := x.Symbol.ProtoReflect().IsValid()

			if i && !s || !i && s {
				return nil, tracer.Maskf(runtime.QueryObjectConflictError, "intern and symbol must not be used together")
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *post.UpdateI) (*post.UpdateO, error) {
	return w.han.Update(ctx, req)
}
