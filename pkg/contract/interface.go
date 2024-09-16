package contract

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/contract/uvxcontract"
)

type Interface interface {
	Claims() claimscontract.Interface
	UVX() uvxcontract.Interface
	Wait(txn *types.Transaction, max time.Duration) error
}
