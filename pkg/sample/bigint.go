package sample

import (
	"math/big"

	"github.com/uvio-network/apiserver/pkg/runtime"
)

func BigInt(a []uint64, b []uint64) []*big.Int {
	var lis []*big.Int

	// Convert the given integers of a to *big.Int while keeping their numerical
	// representation. 1 remains 1 and 5 remains 5.
	for _, x := range a {
		lis = append(lis, big.NewInt(0).SetUint64(x))
	}

	// Convert the given integers of b to *big.Int while translating their
	// numerical value by subtracting from the max uint256. 1 becomes (2^256-1)-1
	// and 5 becomes (2^256-1)-5.
	for _, y := range b {
		lis = append(lis, new(big.Int).Sub(runtime.MaxUint256(), big.NewInt(0).SetUint64(y)))
	}

	return lis
}
