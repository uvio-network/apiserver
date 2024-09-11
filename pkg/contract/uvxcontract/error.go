package uvxcontract

import "github.com/xh3b4sd/tracer"

var BalanceConversionError = &tracer.Error{
	Kind: "BalanceConversionError",
	Desc: "The provided balance was found to be too high, which may indicate that some inputs were misconfigured.",
}
