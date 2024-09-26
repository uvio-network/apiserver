package uvxminthandler

import "github.com/xh3b4sd/tracer"

var WalletNotFoundError = &tracer.Error{
	Kind: "WalletNotFoundError",
	Desc: "The request expects the user to have a wallet. No wallet could be found for the user. Therefore the request failed.",
}
