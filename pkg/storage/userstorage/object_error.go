package userstorage

import (
	"github.com/xh3b4sd/tracer"
)

var UserNameLengthError = &tracer.Error{
	Kind: "UserNameLengthError",
	Desc: "The request expects the user name to have between 2 and 30 characters. The user name was not found to have between 2 and 30 characters. Therefore the request failed.",
}

var UserSubjectEmptyError = &tracer.Error{
	Kind: "UserSubjectEmptyError",
	Desc: "The request expects the user's subject claim not to be empty. The user's subject claim was found to be empty. Therefore the request failed.",
}

var UserSummaryError = &tracer.Error{
	Kind: "UserSummaryError",
	Desc: "The request expects the user's summary to have exactly 5 elements. The user's summary was not found to have 5 elements. Therefore the request failed.",
}
