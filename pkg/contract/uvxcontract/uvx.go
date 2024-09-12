package uvxcontract

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type UVXConfig struct {
	Add common.Address
	Cli *ethclient.Client
	Log logger.Interface
	Opt *bind.TransactOpts
}

type UVX struct {
	bin *UvxContractBinding
	cli *ethclient.Client
	log logger.Interface
	opt *bind.TransactOpts
}

func NewUVX(c UVXConfig) *UVX {
	if len(c.Add) == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Add must not be empty", c)))
	}
	if c.Cli == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cli must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Opt == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Opt must not be empty", c)))
	}

	var err error

	var bin *UvxContractBinding
	{
		bin, err = NewUvxContractBinding(c.Add, c.Cli)
		if err != nil {
			tracer.Panic(err)
		}
	}

	return &UVX{
		bin: bin,
		cli: c.Cli,
		log: c.Log,
		opt: c.Opt,
	}
}

func (u *UVX) Mint(dst string, bal int64) (*types.Transaction, error) {
	var err error

	// In case a balance is provided that appears to account for decimals already,
	// we reject the minting request. The Mint interface states that Mint is
	// responsible for converting decimals according to the provided nominal
	// value. Nobody should ever be able to mint exorbitant amounts of tokens,
	// especially not by accident.
	if bal > 1_000_000 {
		return nil, tracer.Mask(BalanceConversionError)
	}

	var opt *bind.TransactOpts
	{
		opt = &bind.TransactOpts{
			From: u.opt.From,

			// Here we are trying to set some reasonable gas limits, specifically for
			// the EIP-1559 enabled minting transaction.
			//
			//     GasFeeCap is the max gas fee we are willing to pay
			//     GasTipCap is the max priority fee we are willing to pay
			//
			// Below is a testnet transaction providing some real world insight into
			// effective gas usage.
			//
			//     https://sepolia.basescan.org/tx/0x036cf41f0e0187848c0365d91eb368c8ffd589f2794a34caba5cd2609ca8f00a
			//
			// Below is a dune dashboard to show current and historical gas metrics on
			// the Base L2.
			//
			//     https://dune.com/payton/base-l2-gas-price-tracker
			//
			GasFeeCap: big.NewInt(30_000_000), // 0.030 gwei
			GasTipCap: big.NewInt(3_000_000),  // 0.003 gwei

			Signer: u.opt.Signer,
		}
	}

	var txn *types.Transaction
	{
		txn, err = u.bin.Mint(opt, common.HexToAddress(dst), addDec(bal))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u.log.Log(
		context.Background(),
		"level", "debug",
		"message", "submitted UVX minting transaction onchain",
		"signer", u.opt.From.Hex(),
		"address", dst,
		"amount", strconv.FormatInt(bal, 10),
		"transaction", txn.Hash().Hex(),
		"stack", tracer.Stack(err),
	)

	return txn, nil
}

// addDec returns the given balance with 18 additional decimals. 18 decimals are
// what the UVX contract uses. And so what is being returned by addDec is ready
// to be used for onchain transactions.
func addDec(bal int64) *big.Int {
	dec := new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(18), // 10^18
		nil,
	)

	return new(big.Int).Mul(
		big.NewInt(bal),
		dec,
	)
}
