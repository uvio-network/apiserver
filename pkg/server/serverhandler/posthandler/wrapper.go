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
			p := createPublicEmpty(x.Public)

			if p {
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
			p := searchPublicEmpty(x.Public)
			s := searchSymbolEmpty(x.Symbol)

			if !i && !p {
				return nil, tracer.Maskf(runtime.QueryObjectConflictError, "intern and public must not be used together")
			}
			if !i && !s {
				return nil, tracer.Maskf(runtime.QueryObjectConflictError, "intern and symbol must not be used together")
			}
			if !p && !s {
				return nil, tracer.Maskf(runtime.QueryObjectConflictError, "public and symbol must not be used together")
			}
			if i && p && s {
				return nil, tracer.Maskf(runtime.QueryObjectEmptyError, "one of [intern public symbol] must be used")
			}

			if !s && x.Symbol.Vote == "self" && userid.FromContext(ctx) == "" {
				return nil, tracer.Maskf(runtime.UserAuthError, "symbol.vote=self requires valid access token")
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *post.UpdateI) (*post.UpdateO, error) {
	return w.han.Update(ctx, req)
}

func createPublicEmpty(x *post.CreateI_Object_Public) bool {
	return x == nil || (x.Chain == "" && x.Expiry == "" && x.Hash == "" && x.Kind == "" && x.Labels == "" && x.Lifecycle == "" && x.Meta == "" && x.Parent == "" && x.Text == "" && x.Token == "")
}

func searchInternEmpty(x *post.SearchI_Object_Intern) bool {
	return x == nil || (x.Id == "" && x.Owner == "" && x.Tree == "")
}

func searchPublicEmpty(x *post.SearchI_Object_Public) bool {
	return x == nil || (x.Labels == "")
}

func searchSymbolEmpty(x *post.SearchI_Object_Symbol) bool {
	return x == nil || (x.List == "" && x.Time == "" && x.Vote == "")
}
