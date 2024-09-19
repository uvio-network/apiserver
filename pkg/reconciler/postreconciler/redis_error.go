package postreconciler

import (
	"github.com/xh3b4sd/tracer"
)

var ClaimAlreadyExpiredError = &tracer.Error{
	Kind: "ClaimAlreadyExpiredError",
	Desc: "The request expects the claim expiry not to have passed already. The claim expiry was found to have passed already. Therefore the request failed.",
}

var ClaimLifecycleOnchainError = &tracer.Error{
	Kind: "ClaimLifecycleOnchainError",
	Desc: "The request expects the claim to be pending in order to be deleted. The claim was found to be confirmed onchain already. Therefore the request failed.",
}

var ClaimUpdateHashError = &tracer.Error{
	Kind: "ClaimUpdateHashError",
	Desc: "The request expects the claim hash to be updated only once. The claim hash was found to be updated already. Therefore the request failed.",
}

var ClaimUpdateKindError = &tracer.Error{
	Kind: "ClaimUpdateKindError",
	Desc: "The request expects the post ID to reference a claim. The post ID was not found to reference a claim. Therefore the request failed.",
}

var PostOwnerVoteError = &tracer.Error{
	Kind: "PostOwnerVoteError",
	Desc: "The request expects the calling user to have participated in the referenced market. The calling user was not found to have participated in the referenced market. Therefore the request failed.",
}

var PostParentKindError = &tracer.Error{
	Kind: "PostParentKindError",
	Desc: "The request expects the post's parent to reference a claim. The post's parent was not found to reference a claim. Therefore the request failed.",
}
