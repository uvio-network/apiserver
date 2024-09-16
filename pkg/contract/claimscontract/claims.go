package claimscontract

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type ClaimsConfig struct {
	Add common.Address
	Cli *ethclient.Client
	Log logger.Interface
	Opt *bind.TransactOpts
}

type Claims struct {
	bin *ClaimsContractBinding
	cli *ethclient.Client
	log logger.Interface
	opt *bind.TransactOpts
}

func NewClaims(c ClaimsConfig) *Claims {
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

	var bin *ClaimsContractBinding
	{
		bin, err = NewClaimsContractBinding(c.Add, c.Cli)
		if err != nil {
			tracer.Panic(err)
		}
	}

	return &Claims{
		bin: bin,
		cli: c.Cli,
		log: c.Log,
		opt: c.Opt,
	}
}

func (u *Claims) CreateResolve(pro objectid.ID, ind []*big.Int, exp time.Time) (*types.Transaction, error) {
	var err error

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
			//     https://sepolia.basescan.org/tx/TODO
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
		txn, err = u.bin.CreateResolve(opt, big.NewInt(pro.Int()), ind, uint64(exp.Unix()))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	u.log.Log(
		context.Background(),
		"level", "debug",
		"message", "submitted Claims.createResolve transaction onchain",
		"signer", u.opt.From.Hex(),
		"propose", pro.String(),
		"expiry", exp.String(),
		"transaction", txn.Hash().Hex(),
	)

	return txn, nil
}

func (u *Claims) ExistsResolve(pro objectid.ID) (bool, error) {
	var err error

	var res *big.Int
	{
		_, res, err = u.bin.SearchExpired(nil, big.NewInt(pro.Int()))
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return res.Uint64() != 0, nil
}

func (u *Claims) SearchIndices(pro objectid.ID) (uint64, uint64, error) {
	var err error

	var lef *big.Int
	var rig *big.Int
	{
		lef, _, _, _, _, _, _, rig, err = u.bin.SearchIndices(nil, big.NewInt(pro.Int()))
		if err != nil {
			return 0, 0, tracer.Mask(err)
		}
	}

	return lef.Uint64(), rig.Uint64(), nil
}

func (u *Claims) SearchSamples(pro objectid.ID, lef uint64, rig uint64) ([]common.Address, error) {
	var err error

	var add []common.Address
	{
		add, err = u.bin.SearchSamples(nil, big.NewInt(pro.Int()), big.NewInt(0).SetUint64(lef), big.NewInt(0).SetUint64(rig))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return add, nil
}
