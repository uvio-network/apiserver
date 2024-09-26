package uvxcontract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Interface interface {
	// Balance returns the current token balance of the given address.
	Balance(add string) (*big.Int, error)

	// Client is the underlying go-ethereum client interacting with the configured
	// RPC.
	Client() *ethclient.Client

	// Mint generates a signed transaction for minting UVX tokens to the given
	// destination address, to the extend of the given balance amount. Note that
	// bal must be provided as a nominal value without additional decimal
	// conversion. The following example would mint 100 UVX tokens to 0x1234.
	//
	//     Mint("0x1234", 100)
	//
	Mint(dst string, bal int64) (*types.Transaction, error)
}
