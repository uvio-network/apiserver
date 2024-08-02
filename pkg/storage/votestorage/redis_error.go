package votestorage

import (
	"github.com/xh3b4sd/tracer"
)

var VoteObjectNotFoundError = &tracer.Error{
	Kind: "VoteObjectNotFoundError",
	Desc: "The request expects a vote object to exist. The vote object was not found to exist. Therefore the request failed.",
}
