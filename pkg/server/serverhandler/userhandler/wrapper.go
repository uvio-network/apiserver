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
			i := searchInternEmpty(x.Intern)
			p := searchPublicEmpty(x.Public)

			if !i && !p {
				return nil, tracer.Maskf(runtime.QueryObjectConflictError, "intern and public must not be used together")
			}
			if i && p {
				return nil, tracer.Maskf(runtime.QueryObjectEmptyError, "one of [intern public] must be used")
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *user.UpdateI) (*user.UpdateO, error) {
	return w.han.Update(ctx, req)
}

func searchInternEmpty(x *user.SearchI_Object_Intern) bool {
	return x == nil || (x.Id == "")
}

func searchPublicEmpty(x *user.SearchI_Object_Public) bool {
	return x == nil || (x.Name == "")
}
