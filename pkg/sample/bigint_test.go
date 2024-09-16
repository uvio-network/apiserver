package sample

import (
	"fmt"
	"math/big"
	"slices"
	"testing"
)

func Test_Sample_BigInt(t *testing.T) {
	testCases := []struct {
		lef []uint64
		rig []uint64
		res []string
	}{
		// Case 000
		{
			lef: []uint64{},
			rig: []uint64{},
			res: nil,
		},
		// Case 001
		{
			lef: []uint64{1},
			rig: []uint64{3},
			res: []string{
				"1",
				"115792089237316195423570985008687907853269984665640564039457584007913129639932",
			},
		},
		// Case 002
		{
			lef: []uint64{0, 3, 5},
			rig: []uint64{1, 4, 2},
			res: []string{
				"0",
				"3",
				"5",
				"115792089237316195423570985008687907853269984665640564039457584007913129639934",
				"115792089237316195423570985008687907853269984665640564039457584007913129639931",
				"115792089237316195423570985008687907853269984665640564039457584007913129639933",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			res := bigStr(BigInt(tc.lef, tc.rig))

			if !slices.Equal(res, tc.res) {
				t.Fatalf("expected %#v got %#v", tc.res, res)
			}
		})
	}
}

func bigStr(inp []*big.Int) []string {
	var out []string

	for _, x := range inp {
		out = append(out, x.String())
	}

	return out
}
