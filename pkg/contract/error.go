package contract

import "github.com/xh3b4sd/tracer"

var MaxWaitError = &tracer.Error{
	Kind: "MaxWaitError",
	Desc: "The maximum amount of time to wait passed, returning this error.",
}

var TransactionFailedError = &tracer.Error{
	Kind: "TransactionFailedError",
	Desc: "The transaction was mined, but failed with a status code other than 1, which implies that the transaction reverted.",
}
