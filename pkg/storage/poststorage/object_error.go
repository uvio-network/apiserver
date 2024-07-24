package poststorage

import (
	"github.com/xh3b4sd/tracer"
)

var ClaimLifecycleInvalidError = &tracer.Error{
	Kind: "ClaimLifecycleInvalidError",
	Desc: "The request expects the claim lifecycle to be one of [propose resolve dispute nullify]. The claim lifecycle was not found to be one of those values. Therefore the request failed.",
}

var ClaimParentEmptyError = &tracer.Error{
	Kind: "ClaimParentEmptyError",
	Desc: "The request expects the claim parent not to be empty. The claim parent was found to be empty. Therefore the request failed.",
}

var ClaimParentInvalidError = &tracer.Error{
	Kind: "ClaimParentInvalidError",
	Desc: "The request expects the claim parent to be empty. The claim parent was not found to be empty. Therefore the request failed.",
}

var CommentLifecycleInvalidError = &tracer.Error{
	Kind: "CommentLifecycleInvalidError",
	Desc: "The request expects the comment lifecycle to be empty. The comment lifecycle was not found to be empty. Therefore the request failed.",
}

var CommentParentEmptyError = &tracer.Error{
	Kind: "CommentParentEmptyError",
	Desc: "The request expects the comment parent not to be empty. The comment parent was found to be empty. Therefore the request failed.",
}

var PostExpiryEmptyError = &tracer.Error{
	Kind: "PostExpiryEmptyError",
	Desc: "The request expects the post expiry not to be empty. The post expiry was found to be empty. Therefore the request failed.",
}

var PostExpiryPastError = &tracer.Error{
	Kind: "PostExpiryPastError",
	Desc: "The request expects the post expiry not to be in the past. The post expiry was found to be in the past. Therefore the request failed.",
}

var PostKindInvalidError = &tracer.Error{
	Kind: "PostKindInvalidError",
	Desc: "The request expects the post kind to be one of [claim comment]. The post kind was not found to be one of those values. Therefore the request failed.",
}

var PostStakeInvalidError = &tracer.Error{
	Kind: "PostStakeInvalidError",
	Desc: "The request expects the post stake to be a positive number. The post stake was not found to be a positive number. Therefore the request failed.",
}

var PostTextEmptyError = &tracer.Error{
	Kind: "PostTextEmptyError",
	Desc: "The request expects the post text not to be empty. The post text was found to be empty. Therefore the request failed.",
}

var PostTokenEmptyError = &tracer.Error{
	Kind: "PostTokenEmptyError",
	Desc: "The request expects the post token not to be empty. The post token was found to be empty. Therefore the request failed.",
}
