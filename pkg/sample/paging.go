package sample

import (
	"math/big"

	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/xh3b4sd/tracer"
)

const (
	// pag is the batch size of a single page that we want to process at once.
	pag = 25
)

// Paging calculates the next index pair of a split integer sequence based on
// the provided paging pointer. The returned indices are the left and right
// handside cursors that can be used continuously until the boolean "end"
// indicates completion of the iteration by flipping true. "cur" must be 0 in
// order to start an iteration cycle. Consecutive calls to Paging must increment
// "rig" of the previous call and provide the incremented result as "cur" for
// the next call. That means "rig" becomes "lef" by executing some out of band
// process.
func Paging(ind [8]*big.Int, cur *big.Int) (*big.Int, *big.Int, bool, error) {
	// Calculating paginated indices works using the indices [1, 2] and [5, 6]
	// according to the Claims contract documentation linked below.
	//
	//     https://github.com/uvio-network/contracts/blob/v0.5.0/contracts/Claims.sol#L1826-L1882
	//

	for i := range ind {
		if ind[i] == nil {
			return nil, nil, false, tracer.Mask(IndexNilError)
		}
		if bigLss(ind[i], big.NewInt(0)) {
			return nil, nil, false, tracer.Mask(IndexOutOfRangeError)
		}
		if bigGrt(ind[i], runtime.MaxUint256()) {
			return nil, nil, false, tracer.Mask(IndexOutOfRangeError)
		}
	}

	{
		if bigEql(ind[0], big.NewInt(0)) && bigEql(ind[7], big.NewInt(0)) {
			return nil, nil, false, tracer.Mask(IndexZeroError)
		}
		if bigLss(cur, big.NewInt(0)) {
			return nil, nil, false, tracer.Mask(IndexOutOfRangeError)
		}
		if bigGrt(cur, runtime.MaxUint256()) {
			return nil, nil, false, tracer.Mask(IndexOutOfRangeError)
		}
	}

	if bigNot(ind[0], big.NewInt(0)) && (bigEql(cur, big.NewInt(0)) || bigLss(cur, ind[2])) {
		var lef *big.Int
		var rig *big.Int
		var end bool

		if bigEql(cur, big.NewInt(0)) {
			lef = ind[1]
		} else {
			lef = cur
		}

		{
			rig = bigMin(big.NewInt(0).Add(lef, big.NewInt(pag-1)), ind[2])
		}

		if bigEql(rig, ind[2]) && bigEql(ind[7], big.NewInt(0)) {
			end = true
		}

		return lef, rig, end, nil
	}

	if bigNot(ind[7], big.NewInt(0)) && (bigEql(cur, big.NewInt(0)) || bigEql(cur, ind[2]) || bigGrt(cur, ind[2])) {
		var lef *big.Int
		var rig *big.Int
		var end bool

		if bigEql(cur, big.NewInt(0)) || bigLss(cur, ind[5]) {
			lef = ind[5]
		} else {
			lef = cur
		}

		{
			rig = bigMin(big.NewInt(0).Add(lef, big.NewInt(pag-1)), ind[6])
		}

		if bigEql(rig, ind[6]) {
			end = true
		}

		return lef, rig, end, nil
	}

	return nil, nil, false, nil
}

func bigEql(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func bigGrt(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == +1
}

func bigLss(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == -1
}

func bigMin(a *big.Int, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}

	return b
}

func bigNot(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) != 0
}
