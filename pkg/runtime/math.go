package runtime

import "math/big"

var (
	// max is the max uint256 used to translate the indices for the false stakers.
	//
	//     115792089237316195423570985008687907853269984665640564039457584007913129639935
	//
	max = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
)

func MaxUint256() *big.Int {
	return max
}
