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
	"github.com/uvio-network/apiserver/pkg/contract/uvxcontract"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Key is the private key signing transactions for contract writes.
	Key string
	Log logger.Interface
	// RPC is the RPC endpoint for network connection.
	RPC string
	// UVX is the deployed UVX contract address.
	UVX string
}

type Contract struct {
	cli *ethclient.Client
	uvx *uvxcontract.UVX
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

	var uvx *uvxcontract.UVX
	{
		uvx = uvxcontract.NewUVX(uvxcontract.UVXConfig{
			Add: common.HexToAddress(c.UVX),
			Cli: cli,
			Log: c.Log,
			Opt: opt,
		})
	}

	return &Contract{
		cli: cli,
		uvx: uvx,
	}
}

func (c *Contract) UVX() uvxcontract.Interface {
	return c.uvx
}

func (c *Contract) Wait(txn *types.Transaction) error {
	var err error

	for {
		// In today's world we can confidently wait for 1 second at first, since
		// most blocks may not be mined much earlier than that 1 second.
		{
			time.Sleep(1 * time.Second)
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
