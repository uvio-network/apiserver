package userstorage

import (
	"github.com/xh3b4sd/tracer"
)

var UserNotFoundError = &tracer.Error{
	Kind: "UserNotFoundError",
	Desc: "The request expects a user object to exist. The user object was not found to exist. Therefore the request failed.",
}

var SubjectClaimMappingError = &tracer.Error{
	Kind: "SubjectClaimMappingError",
	Desc: "The request expects a mapping between external subject claim and internal user ID. No mapping was found. Therefore the request failed.",
}
