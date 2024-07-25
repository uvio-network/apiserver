package userstorage

import (
	"github.com/xh3b4sd/tracer"
)

var UserNameEmptyError = &tracer.Error{
	Kind: "UserNameEmptyError",
	Desc: "The request expects the user name not to be empty. The user name was found to be empty. Therefore the request failed.",
}

var UserNameLengthError = &tracer.Error{
	Kind: "UserNameLengthError",
	Desc: "The request expects the user name to have between 2 and 30 characters. The user name was not found to have between 2 and 30 characters. Therefore the request failed.",
}

var UserSubjectEmptyError = &tracer.Error{
	Kind: "UserSubjectEmptyError",
	Desc: "The request expects the user's subject claim not to be empty. The user's subject claim was found to be empty. Therefore the request failed.",
}
