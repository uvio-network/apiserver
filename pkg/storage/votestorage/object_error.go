package votestorage

import (
	"github.com/xh3b4sd/tracer"
)

var VoteKindInvalidError = &tracer.Error{
	Kind: "VoteKindInvalidError",
	Desc: "The request expects the vote kind to be one of [stake truth]. The vote kind was not found to be one of those values. Therefore the request failed.",
}

var StakeValueInvalidError = &tracer.Error{
	Kind: "StakeValueInvalidError",
	Desc: "The request expects the vote value to be a positive number. The vote value was not found to be a positive number. Therefore the request failed.",
}

var TruthValueInvalidError = &tracer.Error{
	Kind: "TruthValueInvalidError",
	Desc: "The request expects the truth value to be 1. The truth value was not found to be 1. Therefore the request failed.",
}

var VoteOwnerEmptyError = &tracer.Error{
	Kind: "VoteOwnerEmptyError",
	Desc: "The request expects the vote owner not to be empty. The vote owner was found to be empty. Therefore the request failed.",
}
