package claimscontract

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/objectid"
)

type Interface interface {
	V_0_4_0
	V_0_5_0
}

type V_0_4_0 interface {
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
	SearchIndicesDeprecated(objectid.ID) ([]*big.Int, error)

	// SearchSamples returns the registered voters of the given claim, according
	// to the provided cursors. If a propose has not reached the resolution phase
	// yet, then an empty list is returned.
	SearchSamplesDeprecated(objectid.ID, *big.Int, *big.Int) ([]common.Address, error)

	// SearchVotes returns the number of votes that voted true or false
	// respectively. The returned number of votes can only be considered final if
	// the provided claim has already been settled onchain.
	//
	//     inp[0] the ID of the propose or dispute
	//     out[1] the number of votes cast for true
	//     out[2] the number of votes cast for false
	//
	SearchVotesDeprecated(objectid.ID) (int64, int64, error)
}

// V_0_5_0 is the latest interface with version v0.5.0 that will eventually
// supersede all priors.
type V_0_5_0 interface {
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

	//
	SearchHistory(pod objectid.ID, lef *big.Int, rig *big.Int) ([]*big.Int, error)

	// SearchIndices returns the the total amount of stakers on either side of the
	// market.
	//
	//     inp[0] the propose ID to search indices for
	//     out[0] the total amount of first stakers in agreement
	//     out[1] the left handside index for addresses on the agreeing side
	//     out[2] the right handside index for addresses on the agreeing side
	//     out[3] the left handside index of the proposer address
	//     out[4] the right handside index of the proposer address
	//     out[5] the left handside index for addresses on the disagreeing side
	//     out[6] the right handside index for addresses on the disagreeing side
	//     out[7] the total amount of first stakers in disagreement
	//
	SearchIndices(pod objectid.ID) ([8]*big.Int, error)

	// SearchLatest returns the claim ID of the propose or dispute that represents
	// the last resolvable claim of any given claim tree.
	//
	//     inp[0] the propose or dispute of any given claim tree
	//     out[0] the latest claim ID in the given tree
	//
	SearchLatest(pod objectid.ID) (objectid.ID, error)

	// SearchResolve returns several boolean flags of claim states. With this
	// method it is possible to figure out whether the given claim has already
	// been fully settled onchain.
	//
	//     inp[0] the ID of the propose or dispute to verify
	//     inp[1] the bitmap index to lookup, e.g. CLAIM_BALANCE_S for the settled state
	//     out[0] the boolean flag mapped to the provided bitmap index
	//
	SearchResolve(objectid.ID, uint8) (bool, error)

	// SearchResults returns some structural insights that have been relevant to
	// reconciling any given claim towards settlement.
	//
	//     inp[0] the ID of the propose or dispute
	//     out[1] whether the market resolution was valid
	//     out[2] the side of the market used for settlement
	//     out[3] whether the entire claim tree is considered final
	//
	SearchResults(pod objectid.ID) (bool, bool, bool, error)

	// SearchSamples returns the votes cast per voter according to the addresses
	// returned by SearchVoters. Should only be called on finalized claim trees.
	//
	//     inp[0] the ID of the propose or dispute
	//     inp[1] the left handside index for cast votes to lookup
	//     inp[2] the right handside index for cast votes to lookup
	//     out[0] the cast votes, 0=false, 1=true, 2=abstain
	//
	SearchSamples(objectid.ID, *big.Int, *big.Int) ([]uint8, error)

	//
	SearchStakers(pod objectid.ID, lef *big.Int, rig *big.Int) ([]common.Address, error)

	// SearchVoters returns the registered voter addresses of any given claim,
	// according to the provided cursors. Should only be called on claims with
	// confirmed onchain resolves.
	//
	//     inp[0] the ID of the propose or dispute
	//     inp[1] the left handside index for voter addresses to lookup
	//     inp[2] the right handside index for voter addresses to lookup
	//
	SearchVoters(objectid.ID, *big.Int, *big.Int) ([]common.Address, error)

	// UpdateBalance allows anyone to update user balances in order to settle the
	// originally proposed claim.
	//
	//     inp[0] the ID of the claim to settle
	//     inp[1] the maximum amount of users to update during this call
	//
	UpdateBalance(objectid.ID, uint64) (*types.Transaction, error)

	// Version returns the version of this contract binding and the underlying
	// compatible smart contract.
	Version() string
}
