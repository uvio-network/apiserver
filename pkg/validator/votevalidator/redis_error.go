package votevalidator

import (
	"github.com/xh3b4sd/tracer"
)

var StakeLifecycleInvalidError = &tracer.Error{
	Kind: "StakeLifecycleInvalidError",
	Desc: "The request expects the lifecycle of the referenced claim to be one of [adjourn dispute nullify propose]. The lifecycle of the referenced claim was not found to be one of those values. Therefore the request failed.",
}

var TruthLifecycleInvalidError = &tracer.Error{
	Kind: "TruthLifecycleInvalidError",
	Desc: "The request expects the lifecycle of the referenced claim to be one of [resolve]. The lifecycle of the referenced claim was not found to be one of those values. Therefore the request failed.",
}

var VoteClaimNotFoundError = &tracer.Error{
	Kind: "VoteClaimNotFoundError",
	Desc: "The request expects the vote claim to exist. The vote claim was not found to exist. Therefore the request failed.",
}
