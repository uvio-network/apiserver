package claimscontract

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

type Interface interface {
	// BalanceUpdated searches for all the transaction hashes of the transactions
	// that emitted the event BalanceUpdated. The underlying log filter iterates
	// through all blocks since blc. If no matching event could be found the
	// underlying iterator stops at the latest block.
	//
	//     inp[0] the block number to start filtering from
	//     inp[1] the ID of the propose or dispute for which the balances got updated
	//     out[0] the transaction hashes emitting the filtered event, if any
	//
	BalanceUpdated(uint64, objectid.ID) ([]common.Hash, error)

	// Client is the underlying go-ethereum client interacting with the configured
	// RPC.
	Client() *ethclient.Client

	// CreateResolve allows the BOT_ROLE to initiate the market resolution of the
	// given claim onchain, so that the randomly selected users can verify events
	// in the real world.
	//
	//     inp[0] the ID of the propose or dispute for which to create a resolve
	//     inp[1] the randomized indices of staker addresses
	//     inp[2] the expiry of the resolve to create
	//
	CreateResolve(objectid.ID, []*big.Int, time.Time) (*types.Transaction, error)

	// ExistsResolve returns whether an onchain resolve exists for the given claim.
	//
	//     inp[0] the ID of the propose or dispute to check
	//
	ExistsResolve(objectid.ID) (bool, error)

	// ResolveCreated searches for the transaction hash of the transaction that
	// emitted the event ResolveCreated. The underlying log filter iterates
	// through all blocks since blc until the event matching all required fields
	// was found. If no matching event could be found the underlying iterator
	// stops at the latest block.
	//
	//     inp[0] the block number to start filtering from
	//     inp[1] the ID of the propose for which the resolve got created
	//     out[0] the transaction hash emitting the filtered event, if any
	//
	ResolveCreated(uint64, objectid.ID) (common.Hash, error)

	// SearchIndices returns the the total amount of stakers on either side of the
	// market.
	//
	//     inp[0] the propose ID to search indices for
	//     out[0][0] the total amount of first stakers in agreement
	//     out[0][1] the left handside index for addresses on the agreeing side
	//     out[0][2] the right handside index for addresses on the agreeing side
	//     out[0][3] the left handside index of the proposer address
	//     out[0][4] the right handside index of the proposer address
	//     out[0][5] the left handside index for addresses on the disagreeing side
	//     out[0][6] the right handside index for addresses on the disagreeing side
	//     out[0][7] the total amount of first stakers in disagreement
	//
	SearchIndices(objectid.ID) ([]*big.Int, error)

	// SearchResolve returns several boolean flags of claim states. With this
	// method it is possible to figure out whether the given claim has already
	// been fully settled onchain.
	//
	//     inp[0] the ID of the propose or dispute to verify
	//     inp[1] the bitmap index to lookup, e.g. CLAIM_BALANCE_S for the settled state
	//     out[0] the boolean flag mapped to the provided bitmap index
	//
	SearchResolve(objectid.ID, uint8) (bool, error)

	// SearchSamples returns the registered voters of the given claim, according
	// to the provided cursors. If a propose has not reached the resolution phase
	// yet, then an empty list is returned.
	SearchSamples(objectid.ID, *big.Int, *big.Int) ([]common.Address, error)

	// SearchVotes returns the number of votes that voted true or false
	// respectively. The returned number of votes can only be considered final if
	// the provided claim has already been settled onchain.
	//
	//     inp[0] the ID of the propose or dispute
	//     out[1] the number of votes cast for true
	//     out[2] the number of votes cast for false
	//
	SearchVotes(objectid.ID) (int64, int64, error)

	// UpdateBalance allows anyone to update user balances in order to settle the
	// originally proposed claim.
	//
	//     inp[0] the ID of the claim to settle
	//     inp[1] the maximum amount of users to update during this call
	//
	UpdateBalance(objectid.ID, uint64) (*types.Transaction, error)
}
