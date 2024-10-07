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

type ClaimsConfigV050 struct {
	Add common.Address
	Cli *ethclient.Client
	Log logger.Interface
	Opt *bind.TransactOpts
}

type ClaimsV050 struct {
	bin *ClaimsContractBindingV050
	cli *ethclient.Client
	log logger.Interface
	opt *bind.TransactOpts
}

func NewClaimsV050(c ClaimsConfigV050) *ClaimsV050 {
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

	var bin *ClaimsContractBindingV050
	{
		bin, err = NewClaimsContractBindingV050(c.Add, c.Cli)
		if err != nil {
			tracer.Panic(err)
		}
	}

	return &ClaimsV050{
		bin: bin,
		cli: c.Cli,
		log: c.Log,
		opt: c.Opt,
	}
}

func (c *ClaimsV050) BalanceUpdated(blc uint64, pod objectid.ID) ([]common.Hash, error) {
	var err error

	var ite *ClaimsContractBindingV050BalanceUpdatedIterator
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

		if ite.Event.Pod.Int64() != pod.Int() {
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

func (c *ClaimsV050) CreateResolve(pod objectid.ID, ind []*big.Int, exp time.Time) (*types.Transaction, error) {
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
			//     https://sepolia.basescan.org/tx/0x89688b4baf0efa24ff3fa56b4b01aab04ce303934f0747e38c06adc9c1bea099
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

func (c *ClaimsV050) ExistsResolve(pod objectid.ID) (bool, error) {
	var err error

	var exp *big.Int
	{
		_, exp, err = c.bin.SearchExpired(nil, big.NewInt(pod.Int()))
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return exp.Uint64() != 0, nil
}

func (c *ClaimsV050) ResolveCreated(blc uint64, pod objectid.ID) (common.Hash, error) {
	var err error

	var ite *ClaimsContractBindingV050ResolveCreatedIterator
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

		if ite.Event.Pod.Int64() != pod.Int() {
			continue
		}

		{
			return ite.Event.Raw.TxHash, nil
		}
	}

	return common.Hash{}, nil
}

func (c *ClaimsV050) SearchHistory(pod objectid.ID, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	var err error

	var his []*big.Int
	{
		his, err = c.bin.SearchHistory(nil, big.NewInt(pod.Int()), lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return his, nil
}

func (c *ClaimsV050) SearchIndices(pod objectid.ID) ([8]*big.Int, error) {
	var err error

	var ind [8]*big.Int
	{
		ind, err = c.bin.SearchIndices(nil, big.NewInt(pod.Int()))
		if err != nil {
			return [8]*big.Int{}, tracer.Mask(err)
		}
	}

	return ind, nil
}

func (c *ClaimsV050) SearchResolve(pod objectid.ID, ind uint8) (bool, error) {
	var err error

	var flg bool
	{
		flg, err = c.bin.SearchResolve(nil, big.NewInt(pod.Int()), ind)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return flg, nil
}

func (c *ClaimsV050) SearchResults(pod objectid.ID) (bool, bool, bool, error) {
	var err error

	var val bool
	var sid bool
	var fin bool
	{
		_, val, sid, fin, err = c.bin.SearchResults(nil, big.NewInt(pod.Int()))
		if err != nil {
			return false, false, false, tracer.Mask(err)
		}
	}

	return val, sid, fin, nil
}

func (c *ClaimsV050) SearchSamples(pod objectid.ID, lef *big.Int, rig *big.Int) ([]uint8, error) {
	var err error

	var sam []uint8
	{
		sam, err = c.bin.SearchSamples(nil, big.NewInt(pod.Int()), lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return sam, nil
}

func (c *ClaimsV050) SearchStakers(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var err error

	var sta []common.Address
	{
		sta, err = c.bin.SearchStakers(nil, big.NewInt(pod.Int()), lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return sta, nil
}

func (c *ClaimsV050) SearchVoters(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var err error

	var vot []common.Address
	{
		vot, err = c.bin.SearchVoters(nil, big.NewInt(pod.Int()), lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return vot, nil
}

func (c *ClaimsV050) UpdateBalance(pod objectid.ID, max uint64) (*types.Transaction, error) {
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
