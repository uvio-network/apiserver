package postreconciler

import (
	"github.com/xh3b4sd/tracer"
)

var ClaimAlreadyExpiredError = &tracer.Error{
	Kind: "ClaimAlreadyExpiredError",
	Desc: "The request expects the claim expiry not to have passed already. The claim expiry was found to have passed already. Therefore the request failed.",
}

var ClaimExpiryPastError = &tracer.Error{
	Kind: "ClaimExpiryPastError",
	Desc: "The request expects the claim expiry not to be in the past. The claim expiry was found to be in the past. Therefore the request failed.",
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

var DisputeChallengeError = &tracer.Error{
	Kind: "DisputeChallengeError",
	Desc: "The request expects the disputed resolve to be within its challenge window. The disputed resolve was not found to be within its challenge window anymore. Therefore the request failed.",
}

var DisputeContractError = &tracer.Error{
	Kind: "DisputeContractError",
	Desc: "The request expects the dispute and its referenced resolve to be facilitated by the same smart contract onchain. Either the chain ID, the smart contract address or the token symbol were found to be different. Therefore the request failed.",
}

var DisputeLifecycleError = &tracer.Error{
	Kind: "DisputeLifecycleError",
	Desc: "The request expects disputes to be created on claims with lifecycle phase resolve. The dispute's referenced parent was not found to be of lifecycle phase resolve. Therefore the request failed.",
}

var DisputeLimitError = &tracer.Error{
	Kind: "DisputeLimitError",
	Desc: "The request expects no more than 2 disputes to be created on the originally proposed claim. The originally proposed claim was found to have been disputed twice already. Therefore the request failed.",
}

var DisputeResolutionError = &tracer.Error{
	Kind: "DisputeResolutionError",
	Desc: "The request expects disputes to be created on claims with valid resolutions. The dispute's referenced parent was not found to have a valid resolution. Therefore the request failed.",
}

var MarketParticipationError = &tracer.Error{
	Kind: "MarketParticipationError",
	Desc: "The request expects the calling user to have participated in the referenced market. The calling user was not found to have participated in the referenced market. Therefore the request failed.",
}

var ParentKindError = &tracer.Error{
	Kind: "ParentKindError",
	Desc: "The request expects the post's parent to reference a claim. The post's parent was not found to reference a claim. Therefore the request failed.",
}

var ParentPendingError = &tracer.Error{
	Kind: "ParentPendingError",
	Desc: "The request expects the post's parent to be confirmed onchain. The post's parent was not found to be confirmed onchain. Therefore the request failed.",
}
