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
			if createPublicEmpty(x.Public) {
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
			i := searchInternEmpty(x.Intern)
			s := searchSymbolEmpty(x.Symbol)

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

func createPublicEmpty(x *post.CreateI_Object_Public) bool {
	return x == nil || (x.Expiry == "" && x.Kind == "" && x.Lifecycle == "" && x.Option == "" && x.Parent == "" && x.Stake == "" && x.Text == "" && x.Token == "")
}

func searchInternEmpty(x *post.SearchI_Object_Intern) bool {
	return x == nil || (x.Id == "" && x.Owner == "")
}

func searchSymbolEmpty(x *post.SearchI_Object_Symbol) bool {
	return x == nil || (x.List == "" && x.Time == "" && x.Tree == "")
}
