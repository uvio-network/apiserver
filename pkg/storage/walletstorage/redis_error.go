package walletstorage

import (
	"github.com/xh3b4sd/tracer"
)

var WalletObjectNotFoundError = &tracer.Error{
	Kind: "WalletObjectNotFoundError",
	Desc: "The request expects a wallet object to exist. The wallet object was not found to exist. Therefore the request failed.",
}
