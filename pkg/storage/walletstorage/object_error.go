package walletstorage

import (
	"github.com/xh3b4sd/tracer"
)

var WalletAddressEmptyError = &tracer.Error{
	Kind: "WalletAddressEmptyError",
	Desc: "The request expects the wallet address not to be empty. The wallet address was found to be empty. Therefore the request failed.",
}

var WalletAddressFormatError = &tracer.Error{
	Kind: "WalletAddressFormatError",
	Desc: "The request expects the wallet address to be in hex format of length 42 including 0x prefix. The wallet address was not found to comply with that format. Therefore the request failed.",
}

var WalletDescriptionLimitError = &tracer.Error{
	Kind: "WalletDescriptionLimitError",
	Desc: "The request expects the wallet description not to have more than 100 characters. The wallet description was found to have more than 100 characters. Therefore the request failed.",
}

var WalletKindInvalidError = &tracer.Error{
	Kind: "WalletKindInvalidError",
	Desc: "The request expects the wallet kind to be one of [contract signer]. The wallet kind was not found to be one of those values. Therefore the request failed.",
}

var WalletOwnerEmptyError = &tracer.Error{
	Kind: "WalletOwnerEmptyError",
	Desc: "The request expects the wallet owner not to be empty. The wallet owner was found to be empty. Therefore the request failed.",
}

var WalletProviderInvalidError = &tracer.Error{
	Kind: "WalletProviderInvalidError",
	Desc: "The request expects the wallet provider to be one of [biconomy external privy]. The wallet provider was not found to be one of those values. Therefore the request failed.",
}
