package claimscontract

import (
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

const (
	// CLAIM_BALANCE_S is the bitmap index to call SearchResolve with. This index
	// is mapped to a boolean that tracks whether claims got already fully settled
	// onchain.
	CLAIM_BALANCE_S uint8 = 1
)

type ClaimsConfig struct {
	Cli *ethclient.Client
	Con *Contract
	Log logger.Interface
	Opt *bind.TransactOpts
}

type Claims struct {
	cli *ethclient.Client
	con *Contract
	log logger.Interface
	opt *bind.TransactOpts
	v40 *ClaimsV040
	v50 *ClaimsV050
}

func NewClaims(c ClaimsConfig) *Claims {
	if c.Cli == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cli must not be empty", c)))
	}
	if c.Con == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Con must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Opt == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Opt must not be empty", c)))
	}

	var v40 *ClaimsV040
	if c.Con.Version == "v0.4.0" {
		v40 = NewClaimsV040(ClaimsConfigV040{
			Add: c.Con.Address,
			Cli: c.Cli,
			Log: c.Log,
			Opt: c.Opt,
		})
	}

	var v50 *ClaimsV050
	if c.Con.Version == "v0.5.0" {
		v50 = NewClaimsV050(ClaimsConfigV050{
			Add: c.Con.Address,
			Cli: c.Cli,
			Log: c.Log,
			Opt: c.Opt,
		})
	}

	return &Claims{
		cli: c.Cli,
		con: c.Con,
		log: c.Log,
		opt: c.Opt,
		v40: v40,
		v50: v50,
	}
}

func (c *Claims) BalanceUpdated(blc uint64, pod objectid.ID) ([]common.Hash, error) {
	var err error

	var hsh []common.Hash

	if c.v40 != nil {
		hsh, err = c.v40.BalanceUpdated(blc, pod)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		hsh, err = c.v50.BalanceUpdated(blc, pod)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return hsh, nil
}

func (c *Claims) Client() *ethclient.Client {
	return c.cli
}

func (c *Claims) CreateResolve(pod objectid.ID, ind []*big.Int, exp time.Time) (*types.Transaction, error) {
	var err error

	var txn *types.Transaction

	if c.v40 != nil {
		txn, err = c.v40.CreateResolve(pod, ind, exp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		txn, err = c.v50.CreateResolve(pod, ind, exp)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return txn, nil
}

func (c *Claims) ExistsResolve(pod objectid.ID) (bool, error) {
	var err error

	var exi bool

	if c.v40 != nil {
		exi, err = c.v40.ExistsResolve(pod)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		exi, err = c.v50.ExistsResolve(pod)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return exi, nil
}

func (c *Claims) ResolveCreated(blc uint64, pod objectid.ID) (common.Hash, error) {
	var err error

	var hsh common.Hash

	if c.v40 != nil {
		hsh, err = c.v40.ResolveCreated(blc, pod)
		if err != nil {
			return common.Hash{}, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		hsh, err = c.v50.ResolveCreated(blc, pod)
		if err != nil {
			return common.Hash{}, tracer.Mask(err)
		}
	}

	return hsh, nil
}

func (c *Claims) SearchHistory(pod objectid.ID, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	var err error

	var his []*big.Int

	if c.v40 != nil {
		his, err = c.v40.SearchHistory(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		his, err = c.v50.SearchHistory(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return his, nil
}

func (c *Claims) SearchIndices(pod objectid.ID) ([8]*big.Int, error) {
	var err error

	var ind [8]*big.Int

	if c.v50 != nil {
		ind, err = c.v50.SearchIndices(pod)
		if err != nil {
			return [8]*big.Int{}, tracer.Mask(err)
		}
	}

	return ind, nil
}

func (c *Claims) SearchIndicesDeprecated(pod objectid.ID) ([]*big.Int, error) {
	var err error

	var ind []*big.Int

	if c.v40 != nil {
		ind, err = c.v40.SearchIndices(pod)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return ind, nil
}

func (c *Claims) SearchResolve(pod objectid.ID, ind uint8) (bool, error) {
	var err error

	var flg bool

	if c.v40 != nil {
		flg, err = c.v40.SearchResolve(pod, ind)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		flg, err = c.v50.SearchResolve(pod, ind)
		if err != nil {
			return false, tracer.Mask(err)
		}
	}

	return flg, nil
}

func (c *Claims) SearchSamples(pod objectid.ID, lef *big.Int, rig *big.Int) ([]uint8, error) {
	var err error

	var sam []uint8
	if c.v50 != nil {
		sam, err = c.v50.SearchSamples(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return sam, nil
}

func (c *Claims) SearchSamplesDeprecated(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var err error

	var add []common.Address
	if c.v40 != nil {
		add, err = c.v40.SearchSamples(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return add, nil
}

func (c *Claims) SearchStakers(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var err error

	var sta []common.Address

	if c.v40 != nil {
		sta, err = c.v40.SearchStakers(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		sta, err = c.v50.SearchStakers(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return sta, nil
}

func (c *Claims) SearchVotesDeprecated(pod objectid.ID) (int64, int64, error) {
	var err error

	var yay int64
	var nah int64
	if c.v40 != nil {
		yay, nah, err = c.v40.SearchVotes(pod)
		if err != nil {
			return 0, 0, tracer.Mask(err)
		}
	}

	return yay, nah, nil
}

func (c *Claims) SearchVoters(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var err error

	var vot []common.Address
	if c.v50 != nil {
		vot, err = c.v50.SearchVoters(pod, lef, rig)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return vot, nil
}

func (c *Claims) UpdateBalance(pod objectid.ID, max uint64) (*types.Transaction, error) {
	var err error

	var txn *types.Transaction

	if c.v40 != nil {
		txn, err = c.v40.UpdateBalance(pod, max)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if c.v50 != nil {
		txn, err = c.v50.UpdateBalance(pod, max)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return txn, nil
}

func (c *Claims) Version() string {
	return c.con.Version
}
