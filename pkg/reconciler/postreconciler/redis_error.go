package postreconciler

import (
	"github.com/xh3b4sd/tracer"
)

var PostOwnerVoteError = &tracer.Error{
	Kind: "PostOwnerVoteError",
	Desc: "The request expects the calling user to have participated in the referenced market. The calling user was not found to have participated in the referenced market. Therefore the request failed.",
}

var PostParentKindError = &tracer.Error{
	Kind: "PostParentKindError",
	Desc: "The request expects the post's parent to reference a claim. The post's parent was not found to reference a claim. Therefore the request failed.",
}
