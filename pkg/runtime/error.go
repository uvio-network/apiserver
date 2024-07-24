package runtime

import (
	"github.com/xh3b4sd/tracer"
)

var ExecutionFailedError = &tracer.Error{
	Kind: "ExecutionFailedError",
	Desc: "This internal error implies a severe malfunction of the system.",
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
