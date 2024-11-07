package runtime

import (
	"github.com/xh3b4sd/tracer"
)

var ExecutionFailedError = &tracer.Error{
	Kind: "ExecutionFailedError",
	Desc: "This internal error implies a severe malfunction of the system.",
}

var QueryObjectConflictError = &tracer.Error{
	Kind: "QueryObjectConflictError",
	Desc: "The request expects the query object to define a single intent. The query object was found to define multiple intents. Therefore the request failed.",
}

var QueryObjectEmptyError = &tracer.Error{
	Kind: "QueryObjectEmptyError",
	Desc: "The request expects the query object not to be empty. The query object was found to be empty. Therefore the request failed.",
}

var QueryObjectInvalidError = &tracer.Error{
	Kind: "QueryObjectInvalidError",
	Desc: "The request expects the query object not be valid. The query object was not found to be valid. Therefore the request failed.",
}

var QueryObjectLimitError = &tracer.Error{
	Kind: "QueryObjectLimitError",
	Desc: "The request expects an upper limit of 100 query objects per call. The upper limit of 100 query objects per call was found. Therefore the request failed.",
}

var QueryPagingInvalidError = &tracer.Error{
	Kind: "QueryPagingInvalidError",
	Desc: "The request expects the query paging to consist of two positive integers. At least one of the paging pointers was found to not be a number. Therefore the request failed.",
}

var QueryPagingMissingError = &tracer.Error{
	Kind: "QueryPagingMissingError",
	Desc: "The request expects the query paging to consist of two positive integers. At least one of the paging pointers was found to be missing. Therefore the request failed.",
}

var QueryPagingNegativeError = &tracer.Error{
	Kind: "QueryPagingNegativeError",
	Desc: "The request expects the query paging to consist of two positive integers. One of the paging pointers was found to be negative. Therefore the request failed.",
}

var QueryPagingPageError = &tracer.Error{
	Kind: "QueryPagingPageError",
	Desc: "The request expects the query paging range to ask for at least one, but not more than 1000 objects. The query paging range was found to be outside of those boundary conditions. Therefore the request failed.",
}

var QueryPagingTimeError = &tracer.Error{
	Kind: "QueryPagingTimeError",
	Desc: "The request expects the query paging range to define a difference of at least one second. The query paging range was found to be outside of this boundary condition. Therefore the request failed.",
}

var UserAuthError = &tracer.Error{
	Kind: "UserAuthError",
	Desc: "The request expects a valid access token. No valid access token was found. Therefore the request failed.",
}

var UserNotOwnerError = &tracer.Error{
	Kind: "UserNotOwnerError",
	Desc: "The request expects the calling user to be the owner of the requested resource. The calling user was not found to be the owner of the requested resource. Therefore the request failed.",
}
