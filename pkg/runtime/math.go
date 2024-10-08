package runtime

import (
	"fmt"
	"math/big"

	"github.com/xh3b4sd/tracer"
)

const (
	// max represents the end of the possible uint256 sequence.
	//
	//     uint256 public constant MAX_UINT256 = type(uint256).max;
	//
	max = "115792089237316195423570985008687907853269984665640564039457584007913129639935"

	// MID_UINT256 represents the mid point of the possible uint256 sequence. We
	// use this mid point to identify on which sides our indices are along the
	// integer sequence.
	//
	//     uint256 public constant MID_UINT256 = type(uint256).max / 2;
	//
	mid = "57896044618658097711785492504343953926634992332820282019728792003956564819967"
)

func MaxUint256() *big.Int {
	typ, ok := new(big.Int).SetString(max, 10)
	if !ok {
		tracer.Panic(fmt.Errorf("cannot convert %s to *big.Int", max))
	}

	return typ
}

func MidUint256() *big.Int {
	typ, ok := new(big.Int).SetString(mid, 10)
	if !ok {
		tracer.Panic(fmt.Errorf("cannot convert %s to *big.Int", mid))
	}

	return typ
}
