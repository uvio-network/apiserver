package contract

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/contract/uvxcontract"
)

type Interface interface {
	// UVX returns a Claims contract interface for the given address and the
	// underlying RPC configured.
	Claims(string) claimscontract.Interface

	// Client is the underlying go-ethereum client interacting with the configured
	// RPC.
	Client() *ethclient.Client

	// UVX returns a UVX contract interface for the given address and the
	// underlying RPC configured.
	UVX(string) uvxcontract.Interface

	// Wait allows the caller to wait for the given transaction to be mined
	// onchain, but only for the maximum amount of time provided. Should the
	// provided transaction not be mined successfully within "max", then Wait will
	// return MaxWaitError. Should the provided transaction fail in time, then
	// TransactionFailedError is returned.
	Wait(txn *types.Transaction, max time.Duration) error
}
