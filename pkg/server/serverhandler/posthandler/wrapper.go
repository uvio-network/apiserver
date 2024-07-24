package posthandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/post"
	"github.com/uvio-network/apiserver/pkg/runtime"
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

	// {
	// 	if userid.FromContext(ctx) == "" {
	// 		return nil, tracer.Mask(runtime.UserIDEmptyError)
	// 	}
	// }

	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *post.DeleteI) (*post.DeleteO, error) {
	// {
	// 	if len(req.Object) == 0 {
	// 		return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 	}

	// 	if len(req.Object) > 100 {
	// 		return nil, tracer.Mask(runtime.QueryObjectLimitError)
	// 	}

	// 	for _, x := range req.Object {
	// 		if x == nil {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}
	// }

	// {
	// 	for _, x := range req.Object {
	// 		if x.Intern == nil {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}

	// 	for _, x := range req.Object {
	// 		if x.Intern != nil && x.Intern.Desc == "" {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}
	// }

	// {
	// 	if userid.FromContext(ctx) == "" {
	// 		return nil, tracer.Mask(runtime.UserIDEmptyError)
	// 	}
	// }

	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *post.SearchI) (*post.SearchO, error) {
	// {
	// 	if len(req.Object) == 0 {
	// 		return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 	}

	// 	if len(req.Object) > 100 {
	// 		return nil, tracer.Mask(runtime.QueryObjectLimitError)
	// 	}

	// 	for _, x := range req.Object {
	// 		if x == nil {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}
	// }

	// {
	// 	for _, x := range req.Object {
	// 		if x.Public == nil {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}

	// 	for _, x := range req.Object {
	// 		if x.Public != nil && x.Public.Evnt == "" {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}
	// }

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *post.UpdateI) (*post.UpdateO, error) {
	// {
	// 	if len(req.Object) == 0 {
	// 		return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 	}

	// 	if len(req.Object) > 100 {
	// 		return nil, tracer.Mask(runtime.QueryObjectLimitError)
	// 	}

	// 	for _, x := range req.Object {
	// 		if x == nil {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}
	// }

	// {
	// 	for _, x := range req.Object {
	// 		if x.Intern == nil && x.Symbol == nil && x.Update == nil {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}

	// 	for _, x := range req.Object {
	// 		if x.Intern == nil {
	// 			return nil, tracer.Mask(updateEmptyError)
	// 		}
	// 		if x.Intern != nil && (x.Symbol == nil && x.Update == nil) {
	// 			return nil, tracer.Mask(updateEmptyError)
	// 		}
	// 		if x.Intern != nil && (x.Symbol != nil && x.Update != nil) {
	// 			return nil, tracer.Mask(updateSymbolConflictError)
	// 		}
	// 	}

	// 	for _, x := range req.Object {
	// 		if x.Intern != nil && x.Intern.Desc == "" {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 		if x.Symbol != nil && x.Symbol.Like == "" {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 		if x.Update != nil && len(x.Update) == 0 {
	// 			return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 		}
	// 	}

	// 	for _, x := range req.Object {
	// 		if x.Symbol != nil && x.Symbol.Like != "add" && x.Symbol.Like != "rem" {
	// 			return nil, tracer.Mask(updateSymbolInvalidError)
	// 		}
	// 	}

	// 	for _, x := range req.Object {
	// 		for _, y := range x.Update {
	// 			if y == nil {
	// 				return nil, tracer.Mask(runtime.QueryObjectEmptyError)
	// 			}
	// 		}
	// 	}
	// }

	// {
	// 	if userid.FromContext(ctx) == "" {
	// 		return nil, tracer.Mask(runtime.UserIDEmptyError)
	// 	}
	// }

	return w.han.Update(ctx, req)
}
