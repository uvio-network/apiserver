package votestorage

import (
	"github.com/xh3b4sd/tracer"
)

var StakeValueInvalidError = &tracer.Error{
	Kind: "StakeValueInvalidError",
	Desc: "The request expects the stake value to be a positive number. The stake value was not found to be a positive number. Therefore the request failed.",
}

var TruthValueInvalidError = &tracer.Error{
	Kind: "TruthValueInvalidError",
	Desc: "The request expects the truth value to be 1. The truth value was not found to be 1. Therefore the request failed.",
}

var VoteChainEmptyError = &tracer.Error{
	Kind: "VoteChainEmptyError",
	Desc: "The request expects the vote chain not to be empty. The vote chain was found to be empty. Therefore the request failed.",
}

var VoteHashFormatError = &tracer.Error{
	Kind: "VoteHashFormatError",
	Desc: "The request expects the vote hash to be in hex format including 0x prefix. The vote hash was not found to comply with that format. Therefore the request failed.",
}

var VoteKindInvalidError = &tracer.Error{
	Kind: "VoteKindInvalidError",
	Desc: "The request expects the vote kind to be one of [stake truth]. The vote kind was not found to be one of those values. Therefore the request failed.",
}

var VoteLifecycleInvalidError = &tracer.Error{
	Kind: "VoteLifecycleInvalidError",
	Desc: "The request expects the vote lifecycle to be one of [onchain pending]. The vote lifecycle was not found to be one of those values. Therefore the request failed.",
}

var VoteOwnerEmptyError = &tracer.Error{
	Kind: "VoteOwnerEmptyError",
	Desc: "The request expects the vote owner not to be empty. The vote owner was found to be empty. Therefore the request failed.",
}
