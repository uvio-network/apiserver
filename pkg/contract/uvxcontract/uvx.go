package uvxcontract

import (
	"fmt"
	"math/big"

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
			From:   u.opt.From,
			Signer: u.opt.Signer,
			// TODO set some reasonable gas limits
			//
			//     GasPrice   *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
			//     GasFeeCap  *big.Int // Gas fee cap to use for the 1559 transaction execution (nil = gas price oracle)
			//     GasTipCap  *big.Int // Gas priority fee cap to use for the 1559 transaction execution (nil = gas price oracle)
			//     GasLimit   uint64   // Gas limit to set for the transaction execution (0 = estimate)
			//
		}
	}

	var txn *types.Transaction
	{
		txn, err = u.bin.Mint(opt, common.HexToAddress(dst), addDec(bal))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

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
