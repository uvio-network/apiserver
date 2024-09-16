package claimscontract

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

type Interface interface {
	// CreateResolve initiates the market resolution of the given propose onchain,
	// so that the randomly selected users can verify events in the real world.
	CreateResolve(objectid.ID, []*big.Int, time.Time) (*types.Transaction, error)

	// ExistsResolve returns whether an onchain resolve exists for the given propose.
	ExistsResolve(objectid.ID) (bool, error)

	// SearchIndices returns the the total amount of stakers on either side of the
	// market.
	//
	//     inp[0] the propose ID to search indices for
	//     out[0] the total amount of first stakers in agreement
	//     out[1] the total amount of first stakers in disagreement
	//
	SearchIndices(objectid.ID) (uint64, uint64, error)

	// SearchSamples returns the registered voters of the given propose, according
	// to the provided cursors. If a propose has not reached the resolution phase
	// yet, then an empty list is returned.
	SearchSamples(objectid.ID, uint64, uint64) ([]common.Address, error)
}
