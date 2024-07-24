package poststorage

import (
	"github.com/xh3b4sd/tracer"
)

var PostObjectNotFoundError = &tracer.Error{
	Kind: "PostObjectNotFoundError",
	Desc: "The request expects a post object to exist. The post object was not found to exist. Therefore the request failed.",
}
