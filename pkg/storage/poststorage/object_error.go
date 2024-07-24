package poststorage

import (
	"github.com/xh3b4sd/tracer"
)

var PostTextEmptyError = &tracer.Error{
	Kind: "PostTextEmptyError",
	Desc: "The request expects the post text not to be empty. The post text was found to be empty. Therefore the request failed.",
}

var PostTokenEmptyError = &tracer.Error{
	Kind: "PostTokenEmptyError",
	Desc: "The request expects the post token not to be empty. The post token was found to be empty. Therefore the request failed.",
}
