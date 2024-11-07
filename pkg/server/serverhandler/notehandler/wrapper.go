package notehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/note"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *note.CreateI) (*note.CreateO, error) {
	return w.han.Create(ctx, req)
}

func (w *wrapper) Delete(ctx context.Context, req *note.DeleteI) (*note.DeleteO, error) {
	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *note.SearchI) (*note.SearchO, error) {
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
			p := searchPublicEmpty(x.Public)

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

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *note.UpdateI) (*note.UpdateO, error) {
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
			p := updatePublicEmpty(x.Public)

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

	return w.han.Update(ctx, req)
}

func searchPublicEmpty(x *note.SearchI_Object_Public) bool {
	return x == nil || (x.Kind == "")
}

func updatePublicEmpty(x *note.UpdateI_Object_Public) bool {
	return x == nil || (x.Kind == "" || x.Pointer == "")
}
