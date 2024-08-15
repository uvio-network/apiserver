package votehandler

import (
	"context"

	"github.com/uvio-network/apigocode/pkg/vote"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/context/userid"
	"github.com/xh3b4sd/tracer"
)

type wrapper struct {
	han *Handler
}

func (w *wrapper) Create(ctx context.Context, req *vote.CreateI) (*vote.CreateO, error) {
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

func (w *wrapper) Delete(ctx context.Context, req *vote.DeleteI) (*vote.DeleteO, error) {
	return w.han.Delete(ctx, req)
}

func (w *wrapper) Search(ctx context.Context, req *vote.SearchI) (*vote.SearchO, error) {
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

			if i && p {
				return nil, tracer.Maskf(runtime.QueryObjectEmptyError, "one of [intern public] must be used")
			}

			if !i && x.Intern.Owner == "self" && userid.FromContext(ctx) == "" {
				return nil, tracer.Maskf(runtime.UserAuthError, "intern.owner=self requires valid access token")
			}
		}
	}

	return w.han.Search(ctx, req)
}

func (w *wrapper) Update(ctx context.Context, req *vote.UpdateI) (*vote.UpdateO, error) {
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

func createPublicEmpty(x *vote.CreateI_Object_Public) bool {
	return x == nil || (x.Chain == "" && x.Claim == "" && x.Hash == "" && x.Kind == "" && x.Lifecycle == "" && x.Meta == "" && x.Option == "" && x.Value == "")
}

func searchInternEmpty(x *vote.SearchI_Object_Intern) bool {
	return x == nil || (x.Id == "" && x.Owner == "")
}

func searchPublicEmpty(x *vote.SearchI_Object_Public) bool {
	return x == nil || (x.Claim == "")
}

func updatePublicEmpty(x *vote.UpdateI_Object_Public) bool {
	return x == nil || (x.Hash == "")
}
