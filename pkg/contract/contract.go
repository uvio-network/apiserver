package contract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uvio-network/apiserver/pkg/contract/claimscontract"
	"github.com/uvio-network/apiserver/pkg/contract/uvxcontract"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Key is the private key signing transactions for contract writes.
	Key string
	// Log is a simple logger interface to print system relevant information.
	Log logger.Interface
	// RPC is the RPC endpoint for network connection.
	RPC string
}

type Contract struct {
	cli *ethclient.Client
	log logger.Interface
	opt *bind.TransactOpts
}

func New(c Config) *Contract {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	var err error

	var cli *ethclient.Client
	{
		cli, err = ethclient.Dial(c.RPC)
		if err != nil {
			tracer.Panic(err)
		}
	}

	var cid *big.Int
	{
		cid, err = cli.ChainID(context.Background())
		if err != nil {
			tracer.Panic(err)
		}
	}

	var key *ecdsa.PrivateKey
	{
		key, err = crypto.HexToECDSA(c.Key)
		if err != nil {
			tracer.Panic(err)
		}
	}

	var opt *bind.TransactOpts
	{
		opt, err = bind.NewKeyedTransactorWithChainID(key, cid)
		if err != nil {
			tracer.Panic(err)
		}
	}

	return &Contract{
		cli: cli,
		log: c.Log,
		opt: opt,
	}
}

func (c *Contract) Claims(add string) claimscontract.Interface {
	return claimscontract.NewClaims(claimscontract.ClaimsConfig{
		Add: common.HexToAddress(add),
		Cli: c.cli,
		Log: c.log,
		Opt: c.opt,
	})
}

func (c *Contract) UVX(add string) uvxcontract.Interface {
	return uvxcontract.NewUVX(uvxcontract.UVXConfig{
		Add: common.HexToAddress(add),
		Cli: c.cli,
		Log: c.log,
		Opt: c.opt,
	})
}

func (c *Contract) Wait(txn *types.Transaction, max time.Duration) error {
	var err error

	// Create an outer scope timer for the maximum amount of time to wait during
	// this call.
	var don <-chan time.Time
	{
		don = time.After(max)
	}

	for {
		// In today's world we can confidently wait for 1 second at first, since
		// most blocks may not be mined much earlier than that 1 second.
		var sec <-chan time.Time
		{
			sec = time.After(1 * time.Second)
		}

		select {
		case <-don:
			return tracer.Mask(MaxWaitError)
		case <-sec:
			// fall through
		}

		var rec *types.Receipt
		{
			rec, err = c.cli.TransactionReceipt(context.Background(), txn.Hash())
			if errors.Is(err, ethereum.NotFound) {
				continue
			} else if err != nil {
				return tracer.Mask(err)
			}
		}

		// At this point the transaction was found and there was no error, which
		// means we have a receipt for a mined transaction. What we want to see now
		// is the status field set to 1, which is the specified success status code
		// as per EIP-658.
		//
		//     https://eips.ethereum.org/EIPS/eip-658
		//
		if rec.Status == 1 {
			return nil
		} else {
			return tracer.Mask(TransactionFailedError)
		}
	}
}
