package claimscontract

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	// CLAIM_BALANCE_S is the bitmap index to call SearchResolve with. This index
	// is mapped to a boolean that tracks whether claims got already fully settled
	// onchain.
	CLAIM_BALANCE_S uint8 = 1
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

func (c *Claims) BalanceUpdated(blc uint64, pod uint64) ([]common.Hash, error) {
	var err error

	var ite *ClaimsContractBindingBalanceUpdatedIterator
	{
		ite, err = c.bin.FilterBalanceUpdated(&bind.FilterOpts{Start: blc})
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	{
		defer ite.Close()
	}

	var hsh []common.Hash
	for ite.Next() {
		err := ite.Error()
		if err != nil {
			return nil, tracer.Mask(err)
		}

		if ite.Event.Pod.Uint64() != pod {
			continue
		}

		// Since there are potentially multiple calls to updateBalance, there are
		// potentially multiple events emitted for multiple transaction hashes. And
		// so we collect all of them up to the latest block.
		{
			hsh = append(hsh, ite.Event.Raw.TxHash)
		}
	}

	return hsh, nil
}

func (c *Claims) Client() *ethclient.Client {
	return c.cli
}

func (c *Claims) CreateResolve(pod objectid.ID, ind []*big.Int, exp time.Time) (*types.Transaction, error) {
	var err error

	var opt *bind.TransactOpts
	{
		opt = &bind.TransactOpts{
			From: c.opt.From,

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
			GasFeeCap: big.NewInt(600_000_000), // 0.60 gwei
			GasTipCap: big.NewInt(30_000_000),  // 0.03 gwei

			Signer: c.opt.Signer,
		}
	}

	var txn *types.Transaction
	{
		txn, err = c.bin.CreateResolve(opt, big.NewInt(pod.Int()), ind, uint64(exp.Unix()))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	c.log.Log(
		context.Background(),
		"level", "debug",
		"message", "submitted Claims.createResolve transaction onchain",
		"signer", c.opt.From.Hex(),
		"claim", pod.String(),
		"expiry", exp.String(),
		"transaction", txn.Hash().Hex(),
	)

	return txn, nil
}

func (c *Claims) ExistsResolve(pod objectid.ID) (bool, error) {
	var err error

	var res *big.Int
	{
		_, res, err = c.bin.SearchExpired(nil, big.NewInt(pod.Int()))
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return res.Uint64() != 0, nil
}

func (c *Claims) ResolveCreated(blc uint64, pod uint64) (common.Hash, error) {
	var err error

	var ite *ClaimsContractBindingResolveCreatedIterator
	{
		ite, err = c.bin.FilterResolveCreated(&bind.FilterOpts{Start: blc})
		if err != nil {
			return common.Hash{}, tracer.Mask(err)
		}
	}

	{
		defer ite.Close()
	}

	for ite.Next() {
		err := ite.Error()
		if err != nil {
			return common.Hash{}, tracer.Mask(err)
		}

		if ite.Event.Pod.Uint64() != pod {
			continue
		}

		{
			return ite.Event.Raw.TxHash, nil
		}
	}

	return common.Hash{}, nil
}

func (c *Claims) SearchIndices(pod objectid.ID) ([]*big.Int, error) {
	var err error

	var zer *big.Int
	var one *big.Int
	var two *big.Int
	var thr *big.Int
	var fou *big.Int
	var fiv *big.Int
	var six *big.Int
	var sev *big.Int
	{
		zer, one, two, thr, fou, fiv, six, sev, err = c.bin.SearchIndices(nil, big.NewInt(pod.Int()))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return []*big.Int{zer, one, two, thr, fou, fiv, six, sev}, nil
}

func (c *Claims) SearchResolve(pod objectid.ID, ind uint8) (bool, error) {
	var err error

	var res bool
	{
		res, err = c.bin.SearchResolve(nil, big.NewInt(pod.Int()), ind)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return res, nil
}

func (c *Claims) SearchSamples(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var err error

	var add []common.Address
	{
		add, err = c.bin.SearchSamples(nil, big.NewInt(pod.Int()), lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return add, nil
}

func (c *Claims) SearchVotes(pod objectid.ID) (int64, int64, error) {
	var err error

	var tru *big.Int
	var fls *big.Int
	{
		tru, fls, err = c.bin.SearchVotes(nil, big.NewInt(pod.Int()))
		if err != nil {
			return 0, 0, tracer.Mask(err)
		}
	}

	return tru.Int64(), fls.Int64(), nil
}

func (c *Claims) UpdateBalance(pod objectid.ID, max uint64) (*types.Transaction, error) {
	var err error

	var txn *types.Transaction
	{
		txn, err = c.bin.UpdateBalance(nil, big.NewInt(pod.Int()), big.NewInt(int64(max)))
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	c.log.Log(
		context.Background(),
		"level", "debug",
		"message", "submitted Claims.updateBalance transaction onchain",
		"signer", c.opt.From.Hex(),
		"claim", pod.String(),
		"maximum", strconv.FormatUint(max, 10),
		"transaction", txn.Hash().Hex(),
	)

	return txn, nil
}
