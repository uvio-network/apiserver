package poststorage

import (
	"github.com/xh3b4sd/tracer"
)

var ClaimLabelsEmptyError = &tracer.Error{
	Kind: "ClaimLabelsEmptyError",
	Desc: "The request expects the claim labels not to be empty. The claim labels were found to be empty. Therefore the request failed.",
}

var ClaimLabelsFormatError = &tracer.Error{
	Kind: "ClaimLabelsFormatError",
	Desc: "The request expects the claim labels to be alpha numerical. The claim labels were not found to be alpha numerical. Therefore the request failed.",
}

var ClaimLabelsLimitError = &tracer.Error{
	Kind: "ClaimLabelsLimitError",
	Desc: "The request expects the claim not to have more than four labels. The claim was found to have more than four labels. Therefore the request failed.",
}

var ClaimLabelsUniqueError = &tracer.Error{
	Kind: "ClaimLabelsUniqueError",
	Desc: "The request expects the claim labels to be unique. The claim labels were not found to be unique. Therefore the request failed.",
}

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

var CommentLabelsInvalidError = &tracer.Error{
	Kind: "CommentLabelsInvalidError",
	Desc: "The request expects the comment labels to be empty. The comment labels were not found to be empty. Therefore the request failed.",
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

var PostTextEmptyError = &tracer.Error{
	Kind: "PostTextEmptyError",
	Desc: "The request expects the post text not to be empty. The post text was found to be empty. Therefore the request failed.",
}

var PostTextLengthError = &tracer.Error{
	Kind: "PostTextLengthError",
	Desc: "The request expects the post text to have between 100 and 2500 characters. The post text was not found to have between 100 and 2500 characters. Therefore the request failed.",
}

var PostTokenEmptyError = &tracer.Error{
	Kind: "PostTokenEmptyError",
	Desc: "The request expects the post token not to be empty. The post token was found to be empty. Therefore the request failed.",
}

var PostOwnerEmptyError = &tracer.Error{
	Kind: "PostOwnerEmptyError",
	Desc: "The request expects the post owner not to be empty. The post owner was found to be empty. Therefore the request failed.",
}
