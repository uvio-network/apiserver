package votereconciler

import (
	"github.com/xh3b4sd/tracer"
)

var ClaimAlreadyExpiredError = &tracer.Error{
	Kind: "ClaimAlreadyExpiredError",
	Desc: "The request expects the referenced claim expiry not to have passed already. The referenced claim expiry was found to have passed already. Therefore the request failed.",
}

var ClaimLifecycleInvalidError = &tracer.Error{
	Kind: "ClaimLifecycleInvalidError",
	Desc: "The request expects the lifecycle of the referenced claim to be one of [dispute propose]. The lifecycle of the referenced claim was not found to be one of those values. Therefore the request failed.",
}

var ClaimLifecyclePendingError = &tracer.Error{
	Kind: "ClaimLifecyclePendingError",
	Desc: "The request expects only the proposer to stake reputation while the referenced claim is not confirmed onchain. Somebody other than the proposer was found to stake on the referenced claim. Therefore the request failed.",
}

var StakeLifecyclePendingError = &tracer.Error{
	Kind: "StakeLifecyclePendingError",
	Desc: "The request expects all stakes to be confirmed onchain. A stake was found to exist already in pending state. Therefore the request failed.",
}

var TruthLifecycleInvalidError = &tracer.Error{
	Kind: "TruthLifecycleInvalidError",
	Desc: "The request expects the lifecycle of the referenced claim to be one of [resolve]. The lifecycle of the referenced claim was not found to be one of those values. Therefore the request failed.",
}

var VoteAlreadyExistsError = &tracer.Error{
	Kind: "VoteAlreadyExistsError",
	Desc: "The request expects the user to only vote once on this claim. The user was found to have already voted on this claim. Therefore the request failed.",
}

var VoteClaimNotFoundError = &tracer.Error{
	Kind: "VoteClaimNotFoundError",
	Desc: "The request expects the vote claim to exist. The vote claim was not found to exist. Therefore the request failed.",
}

var VoteLifecycleOnchainError = &tracer.Error{
	Kind: "VoteLifecycleOnchainError",
	Desc: "The request expects the vote to be pending in order to be deleted. The vote was found to be confirmed onchain already. Therefore the request failed.",
}

var VoteLifecyclePendingError = &tracer.Error{
	Kind: "VoteLifecyclePendingError",
	Desc: "The request expects all votes to be confirmed onchain. A vote was found to exist already in pending state. Therefore the request failed.",
}

var VoteUpdateHashError = &tracer.Error{
	Kind: "VoteUpdateHashError",
	Desc: "The request expects the vote hash to be updated only once. The vote hash was found to be updated already. Therefore the request failed.",
}
