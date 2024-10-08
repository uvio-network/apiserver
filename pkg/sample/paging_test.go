package sample

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/uvio-network/apiserver/pkg/runtime"
)

func Test_Sample_Paging_Side_Both(t *testing.T) {
	testCases := []struct {
		ind [8]*big.Int
		cur *big.Int
		lef *big.Int
		rig *big.Int
		end bool
	}{
		// Case 000
		{
			ind: [8]*big.Int{
				big.NewInt(1), // 1 staker in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1), // 1 staker in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(0),
			end: false,
		},
		// Case 001
		{
			ind: [8]*big.Int{
				big.NewInt(1), // 1 staker in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1), // 1 staker in disagreement
			},
			cur: big.NewInt(1),
			lef: runtime.MaxUint256(),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 002
		{
			ind: [8]*big.Int{
				big.NewInt(1), // 1 staker in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
				runtime.MaxUint256(),
				big.NewInt(14), // 14 staker in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(0),
			end: false,
		},
		// Case 003
		{
			ind: [8]*big.Int{
				big.NewInt(1), // 1 staker in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
				runtime.MaxUint256(),
				big.NewInt(14), // 14 staker in disagreement
			},
			cur: big.NewInt(1),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 004
		{
			ind: [8]*big.Int{
				big.NewInt(14), // 14 staker in agreement
				big.NewInt(0),
				big.NewInt(13),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
				runtime.MaxUint256(),
				big.NewInt(14), // 14 staker in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(13),
			end: false,
		},
		// Case 005
		{
			ind: [8]*big.Int{
				big.NewInt(14), // 14 staker in agreement
				big.NewInt(0),
				big.NewInt(13),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
				runtime.MaxUint256(),
				big.NewInt(14), // 14 staker in disagreement
			},
			cur: big.NewInt(14),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
			rig: runtime.MaxUint256(),
			end: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lef, rig, end, err := Paging(tc.ind, tc.cur)
			if err != nil {
				t.Fatal(err)
			}

			if lef.Cmp(tc.lef) != 0 {
				t.Fatalf("expected %#v got %#v", tc.lef.String(), lef.String())
			}
			if rig.Cmp(tc.rig) != 0 {
				t.Fatalf("expected %#v got %#v", tc.rig.String(), rig.String())
			}
			if end != tc.end {
				t.Fatalf("expected %#v got %#v", tc.end, end)
			}
		})
	}
}

func Test_Sample_Paging_Error(t *testing.T) {
	testCases := []struct {
		ind [8]*big.Int
		cur *big.Int
	}{
		// Case 000
		{
			ind: [8]*big.Int{},
			cur: big.NewInt(0),
		},
		// Case 001
		{
			ind: [8]*big.Int{
				nil,
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(0),
		},
		// Case 002
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				nil,
			},
			cur: big.NewInt(0),
		},
		// Case 003
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0),
				big.NewInt(0),
				nil,
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(0),
		},
		// Case 004
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0),
				nil,
				runtime.MidUint256(),
				runtime.MidUint256(),
				nil,
				nil,
				big.NewInt(1),
			},
			cur: big.NewInt(0),
		},
		// Case 005
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(-1),
		},
		// Case 006
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(0).Add(runtime.MaxUint256(), big.NewInt(1)),
		},
		// Case 007
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0).Add(runtime.MaxUint256(), big.NewInt(1)),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(0),
		},
		// Case 008
		{
			ind: [8]*big.Int{
				big.NewInt(0),
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Add(runtime.MaxUint256(), big.NewInt(1)),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(0),
		},
		// Case 9
		{
			ind: [8]*big.Int{
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(1)),
				big.NewInt(0),
				big.NewInt(0),
				big.NewInt(0).Add(runtime.MaxUint256(), big.NewInt(1)),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(1)),
				runtime.MaxUint256(),
				big.NewInt(1),
			},
			cur: big.NewInt(0),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			_, _, _, err := Paging(tc.ind, tc.cur)

			if err == nil {
				t.Fatalf("expected error got nil")
			}
		})
	}
}

func Test_Sample_Paging_Side_False(t *testing.T) {
	testCases := []struct {
		ind [8]*big.Int
		cur *big.Int
		lef *big.Int
		rig *big.Int
		end bool
	}{
		// Case 000
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(1), // 1 staker in disagreement
			},
			cur: big.NewInt(0),
			lef: runtime.MaxUint256(),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 001
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(1)),
				runtime.MaxUint256(),
				big.NewInt(2), // 2 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(1)),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 002
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(23)),
				runtime.MaxUint256(),
				big.NewInt(24), // 24 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(23)),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 003
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(24)),
				runtime.MaxUint256(),
				big.NewInt(25), // 25 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(24)),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 004
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(24)),
				runtime.MaxUint256(),
				big.NewInt(25), // 25 stakers in disagreement
			},
			cur: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(13)),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 005
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(25)),
				runtime.MaxUint256(),
				big.NewInt(26), // 26 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(25)),
			rig: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(1)), // not going over batch size
			end: false,                                                  // there is one more index left to process
		},
		// Case 006
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(25)),
				runtime.MaxUint256(),
				big.NewInt(26), // 26 stakers in disagreement
			},
			cur: runtime.MaxUint256(),
			lef: runtime.MaxUint256(),
			rig: runtime.MaxUint256(),
			end: true,
		},
		// Case 007
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(66)),
				runtime.MaxUint256(),
				big.NewInt(67), // 67 stakers in disagreement
			},
			cur: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(41)),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(41)),
			rig: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(17)),
			end: false,
		},
		// Case 008
		{
			ind: [8]*big.Int{
				big.NewInt(0), // 0 stakers in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(66)),
				runtime.MaxUint256(),
				big.NewInt(67), // 67 stakers in disagreement
			},
			cur: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(16)),
			lef: big.NewInt(0).Sub(runtime.MaxUint256(), big.NewInt(16)),
			rig: runtime.MaxUint256(),
			end: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lef, rig, end, err := Paging(tc.ind, tc.cur)
			if err != nil {
				t.Fatal(err)
			}

			if lef.Cmp(tc.lef) != 0 {
				t.Fatalf("expected %#v got %#v", tc.lef.String(), lef.String())
			}
			if rig.Cmp(tc.rig) != 0 {
				t.Fatalf("expected %#v got %#v", tc.rig.String(), rig.String())
			}
			if end != tc.end {
				t.Fatalf("expected %#v got %#v", tc.end, end)
			}
		})
	}
}

func Test_Sample_Paging_Side_True(t *testing.T) {
	testCases := []struct {
		ind [8]*big.Int
		cur *big.Int
		lef *big.Int
		rig *big.Int
		end bool
	}{
		// Case 000
		{
			ind: [8]*big.Int{
				big.NewInt(1), // 1 staker in agreement
				big.NewInt(0),
				big.NewInt(0),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(0),
			end: true,
		},
		// Case 001
		{
			ind: [8]*big.Int{
				big.NewInt(2), // 2 stakers in agreement
				big.NewInt(0),
				big.NewInt(1),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(1),
			end: true,
		},
		// Case 002
		{
			ind: [8]*big.Int{
				big.NewInt(24), // 24 stakers in agreement
				big.NewInt(0),
				big.NewInt(23),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(23),
			end: true,
		},
		// Case 003
		{
			ind: [8]*big.Int{
				big.NewInt(25), // 25 stakers in agreement
				big.NewInt(0),
				big.NewInt(24),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(24),
			end: true,
		},
		// Case 004
		{
			ind: [8]*big.Int{
				big.NewInt(25), // 25 stakers in agreement
				big.NewInt(0),
				big.NewInt(24),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(11),
			lef: big.NewInt(11),
			rig: big.NewInt(24),
			end: true,
		},
		// Case 005
		{
			ind: [8]*big.Int{
				big.NewInt(26), // 26 stakers in agreement
				big.NewInt(0),
				big.NewInt(25),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(0),
			lef: big.NewInt(0),
			rig: big.NewInt(24), // not going over batch size
			end: false,          // there is one more index left to process
		},
		// Case 006
		{
			ind: [8]*big.Int{
				big.NewInt(26), // 26 stakers in agreement
				big.NewInt(0),
				big.NewInt(25),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(24),
			lef: big.NewInt(24),
			rig: big.NewInt(25),
			end: true,
		},
		// Case 007
		{
			ind: [8]*big.Int{
				big.NewInt(67), // 67 stakers in agreement
				big.NewInt(0),
				big.NewInt(66),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(24),
			lef: big.NewInt(24),
			rig: big.NewInt(48),
			end: false,
		},
		// Case 008
		{
			ind: [8]*big.Int{
				big.NewInt(67), // 67 stakers in agreement
				big.NewInt(0),
				big.NewInt(66),
				runtime.MidUint256(),
				runtime.MidUint256(),
				runtime.MaxUint256(),
				runtime.MaxUint256(),
				big.NewInt(0), // 0 stakers in disagreement
			},
			cur: big.NewInt(49),
			lef: big.NewInt(49),
			rig: big.NewInt(66),
			end: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lef, rig, end, err := Paging(tc.ind, tc.cur)
			if err != nil {
				t.Fatal(err)
			}

			if lef.Cmp(tc.lef) != 0 {
				t.Fatalf("expected %#v got %#v", tc.lef.String(), lef.String())
			}
			if rig.Cmp(tc.rig) != 0 {
				t.Fatalf("expected %#v got %#v", tc.rig.String(), rig.String())
			}
			if end != tc.end {
				t.Fatalf("expected %#v got %#v", tc.end, end)
			}
		})
	}
}
