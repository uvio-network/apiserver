package wallethandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/wallet"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *wallet.CreateI) (*wallet.CreateO, error) {
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

func (w *wrapper) Delete(ctx context.Context, req *wallet.DeleteI) (*wallet.DeleteO, error) {
	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *wallet.SearchI) (*wallet.SearchO, error) {
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

			if i {
				return nil, tracer.Maskf(runtime.QueryObjectEmptyError, "intern must not be empty")
			}

			if !i && x.Intern.Owner != "self" && x.Intern.Owner != string(userid.FromContext(ctx)) {
				return nil, tracer.Mask(runtime.UserNotOwnerError)
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

func (w *wrapper) Update(ctx context.Context, req *wallet.UpdateI) (*wallet.UpdateO, error) {
	return w.han.Update(ctx, req)
}

func createPublicEmpty(x *wallet.CreateI_Object_Public) bool {
	return x == nil || (x.Active == "" && x.Address == "" && x.Description == "" && x.Kind == "" && x.Provider == "")
}

func searchInternEmpty(x *wallet.SearchI_Object_Intern) bool {
	return x == nil || (x.Id == "" && x.Owner == "")
}
