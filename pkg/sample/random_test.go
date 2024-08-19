package sample

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"slices"
	"testing"

	"github.com/uvio-network/apiserver/pkg/generic"
	"github.com/xh3b4sd/tracer"
)

func Test_Sample_Random_crv(t *testing.T) {
	testCases := []struct {
		seq map[uint]int
	}{
		// Case 000
		{
			seq: map[uint]int{
				0: 1,
				1: 1,
				2: 1,
				3: 1,
				4: 1,

				5: 2,
				6: 2,
				7: 2,
				8: 2,
				9: 2,

				10: 3,
				11: 3,
				12: 3,
				13: 3,
				14: 3,

				15: 4,
				16: 4,
				17: 4,
				18: 4,
				19: 4,

				20: 5,
				21: 5,
				22: 5,
				23: 5,
				24: 5,

				25: 5,
				26: 5,
				27: 5,
				28: 5,
				29: 5,

				30: 5,
				31: 5,
				32: 5,
				33: 5,
				34: 5,

				35: 5,
				36: 5,
				37: 5,
				38: 5,
				39: 5,

				40: 5,
				41: 5,
				42: 5,
				43: 5,
				44: 5,

				45: 5,
				46: 5,
				47: 5,
				48: 5,
				49: 5,

				50: 5,
				51: 5,
				52: 5,
				53: 5,
				54: 5,
			},
		},
		// Case 001
		{
			seq: map[uint]int{
				95: 5,
				96: 5,
				97: 5,
				98: 5,
				99: 5,

				100: 5,
				101: 5,
				102: 5,
				103: 5,
				104: 5,

				105: 5,
				106: 5,
				107: 5,
				108: 5,
				109: 5,

				110: 6,
				111: 6,
				112: 6,
				113: 6,
				114: 6,
			},
		},
		// Case 002
		{
			seq: map[uint]int{
				200: 10,
				201: 10,
				202: 10,
				203: 10,
				204: 10,
			},
		},
		// Case 003
		{
			seq: map[uint]int{
				300: 15,
				301: 15,
				302: 15,
				303: 15,
				304: 15,
			},
		},
		// Case 004
		{
			seq: map[uint]int{
				400: 20,
				401: 20,
				402: 20,
				403: 20,
				404: 20,
			},
		},
		// Case 005
		{
			seq: map[uint]int{
				500: 25,
				501: 25,
				502: 25,
				503: 25,
				504: 25,
			},
		},
		// Case 006
		{
			seq: map[uint]int{
				600: 30,
				601: 30,
				602: 30,
				603: 30,
				604: 30,
			},
		},
		// Case 007
		{
			seq: map[uint]int{
				995: 50,
				996: 50,
				997: 50,
				998: 50,
				999: 50,
			},
		},
		// Case 008
		{
			seq: map[uint]int{
				1000:   50,
				1500:   50,
				5000:   50,
				10000:  50,
				100000: 50,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			for k, v := range tc.seq {
				l := crv(k)

				if l != v {
					t.Fatalf("%d: expected %#v got %#v", k, v, l)
				}
			}
		})
	}
}

func Test_Sample_Random_Crypto(t *testing.T) {
	testCases := []struct {
		a uint
		b uint
	}{
		// Case 000
		{
			a: 0,
			b: 0,
		},

		// Case 001
		{
			a: 1,
			b: 1,
		},

		// Case 002
		{
			a: 2,
			b: 2,
		},

		// Case 003
		{
			a: 3,
			b: 3,
		},

		// Case 004
		{
			a: 4,
			b: 4,
		},

		// Case 005
		{
			a: 5,
			b: 5,
		},

		// Case 006
		{
			a: 10,
			b: 10,
		},

		// Case 007
		{
			a: 1,
			b: 0,
		},
		// Case 008
		{
			a: 0,
			b: 1,
		},

		// Case 009
		{
			a: 2,
			b: 0,
		},
		// Case 010
		{
			a: 0,
			b: 2,
		},

		// Case 011
		{
			a: 3,
			b: 0,
		},
		// Case 012
		{
			a: 0,
			b: 3,
		},

		// Case 013
		{
			a: 1,
			b: 2,
		},
		// Case 014
		{
			a: 2,
			b: 1,
		},

		// Case 015
		{
			a: 3,
			b: 1,
		},
		// Case 016
		{
			a: 1,
			b: 3,
		},

		// Case 017
		{
			a: 3,
			b: 2,
		},
		// Case 018
		{
			a: 2,
			b: 3,
		},

		// Case 019
		{
			a: 3,
			b: 4,
		},
		// Case 020
		{
			a: 4,
			b: 3,
		},

		// Case 021
		{
			a: 2,
			b: 5,
		},
		// Case 022
		{
			a: 5,
			b: 2,
		},

		// Case 023
		{
			a: 10,
			b: 0,
		},
		// Case 024
		{
			a: 0,
			b: 10,
		},

		// Case 025
		{
			a: 7,
			b: 5,
		},
		// Case 026
		{
			a: 5,
			b: 7,
		},

		// Case 027
		{
			a: 6,
			b: 1,
		},
		// Case 028
		{
			a: 1,
			b: 6,
		},

		// Case 029
		{
			a: 30,
			b: 5,
		},
		// Case 030
		{
			a: 5,
			b: 30,
		},

		// Case 031
		{
			a: 57,
			b: 2,
		},
		// Case 032
		{
			a: 2,
			b: 57,
		},

		// Case 033
		{
			a: 108,
			b: 20,
		},
		// Case 034
		{
			a: 20,
			b: 108,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			mpx := map[string]int{}
			mpy := map[string]int{}

			for i := 0; i < 1000; i++ {
				var sam *Sample
				{
					sam = New(Config{
						Rea: rand.Reader,
					})
				}

				var x []int
				var y []int
				{
					x, y = sam.Random(tc.a, tc.b)
				}

				{
					mpx[musStr(x)]++
					mpy[musStr(y)]++
				}

				// If we have a positive line on both sides, then the number of random
				// points given to us in return should always be equal.
				if tc.a != 0 && tc.b != 0 && len(x) != len(y) {
					t.Fatalf("expected %#v got %#v", len(x), len(y))
				}

				{
					if generic.Duplicate(x) {
						t.Fatalf("expected x to not contain duplicates")
					}

					if generic.Duplicate(y) {
						t.Fatalf("expected y to not contain duplicates")
					}
				}

				// Below we allocate slices to access items using the generated indices.
				// If we are able to access those items using any given index without
				// creating a runtime panic, then our implementation is safe to the
				// extend that no indices are generated out of bounds.
				{
					slx := make([]int, tc.a)
					for _, i := range x {
						_ = slx[i]
					}

					sly := make([]int, tc.b)
					for _, i := range y {
						_ = sly[i]
					}
				}
			}

			// We somehow want to assert that we generate a variation of random
			// indices. Below we check that there are some minimum amounts of
			// generated permutations, which implies that each of the existing keys in
			// the associated map of observed indices was generated at least once.
			// More keys equals more randomness. And so we compare against the
			// smallest common denominator of the given test case.
			{
				lnx := uint(len(mpx))
				mnx := min(tc.a, tc.b)
				if lnx < mnx {
					t.Fatalf("expected x to generate at least %d different variations, got %d", lnx, mnx)
				}

				lny := uint(len(mpy))
				mny := min(tc.a, tc.b)
				if lny < mny {
					t.Fatalf("expected y to generate at least %d different variations, got %d", lny, mny)
				}
			}
		})
	}
}

func Test_Sample_Random_Fake(t *testing.T) {
	testCases := []struct {
		a uint
		b uint
		s int
		x []int
		y []int
	}{
		// Case 000
		{
			a: 0,
			b: 0,
			s: 0,
			x: nil,
			y: nil,
		},
		// Case 001
		{
			a: 0,
			b: 0,
			s: 3,
			x: nil,
			y: nil,
		},

		// Case 002
		{
			a: 1,
			b: 0,
			s: 0,
			x: []int{
				0,
			},
			y: nil,
		},
		// Case 003
		{
			a: 0,
			b: 1,
			s: 0,
			x: nil,
			y: []int{
				0,
			},
		},

		// Case 004
		{
			a: 1,
			b: 0,
			s: 2376,
			x: []int{
				0,
			},
			y: nil,
		},
		// Case 005
		{
			a: 0,
			b: 1,
			s: 2376,
			x: nil,
			y: []int{
				0,
			},
		},

		// Case 006
		{
			a: 3,
			b: 0,
			s: 0,
			x: []int{
				2,
			},
			y: nil,
		},
		// Case 007
		{
			a: 0,
			b: 3,
			s: 0,
			x: nil,
			y: []int{
				2,
			},
		},

		// Case 008
		{
			a: 3,
			b: 0,
			s: 665117,
			x: []int{
				1,
			},
			y: nil,
		},
		// Case 009
		{
			a: 0,
			b: 3,
			s: 665117,
			x: nil,
			y: []int{
				1,
			},
		},

		// Case 010
		{
			a: 10,
			b: 0,
			s: 0,
			x: []int{
				9,
			},
			y: nil,
		},
		// Case 011
		{
			a: 0,
			b: 10,
			s: 0,
			x: nil,
			y: []int{
				9,
			},
		},

		// Case 012
		{
			a: 10,
			b: 0,
			s: 104017,
			x: []int{
				6,
			},
			y: nil,
		},
		// Case 013
		{
			a: 0,
			b: 10,
			s: 104017,
			x: nil,
			y: []int{
				6,
			},
		},

		// Case 014
		{
			a: 1,
			b: 1,
			s: 0,
			x: []int{
				0,
			},
			y: []int{
				0,
			},
		},
		// Case 015
		{
			a: 1,
			b: 1,
			s: 367425,
			x: []int{
				0,
			},
			y: []int{
				0,
			},
		},

		// Case 016
		{
			a: 1,
			b: 2,
			s: 2376,
			x: []int{
				0,
			},
			y: []int{
				1,
			},
		},
		// Case 017
		{
			a: 2,
			b: 1,
			s: 2376,
			x: []int{
				1,
			},
			y: []int{
				0,
			},
		},

		// Case 018
		{
			a: 3,
			b: 1,
			s: 0,
			x: []int{
				2,
			},
			y: []int{
				0,
			},
		},
		// Case 019
		{
			a: 1,
			b: 3,
			s: 0,
			x: []int{
				0,
			},
			y: []int{
				2,
			},
		},

		// Case 020
		{
			a: 3,
			b: 1,
			s: 665117,
			x: []int{
				1,
			},
			y: []int{
				0,
			},
		},
		// Case 021
		{
			a: 1,
			b: 3,
			s: 665117,
			x: []int{
				0,
			},
			y: []int{
				1,
			},
		},

		// Case 022
		{
			a: 3,
			b: 3,
			s: 0,
			x: []int{
				2,
			},
			y: []int{
				0,
			},
		},
		// Case 023
		{
			a: 3,
			b: 3,
			s: 665117,
			x: []int{
				1,
			},
			y: []int{
				0,
			},
		},

		// Case 024
		{
			a: 3,
			b: 5,
			s: 0,
			x: []int{
				2,
			},
			y: []int{
				0,
			},
		},
		// Case 025
		{
			a: 5,
			b: 3,
			s: 0,
			x: []int{
				4,
			},
			y: []int{
				0,
			},
		},

		// Case 026
		{
			a: 5,
			b: 5,
			s: 29384,
			x: []int{
				3,
				4,
			},
			y: []int{
				0,
				1,
			},
		},
		// Case 027
		{
			a: 10,
			b: 10,
			s: 10197437167,
			x: []int{
				2,
				8,
				6,
			},
			y: []int{
				0,
				1,
				2,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var sam *Sample
			{
				sam = New(Config{
					Rea: newFakeReader(tc.s),
				})
			}

			var x []int
			var y []int
			{
				x, y = sam.Random(tc.a, tc.b)
			}

			{
				if !slices.Equal(x, tc.x) {
					t.Fatalf("expected x %#v got %#v", tc.x, x)
				}

				if !slices.Equal(y, tc.y) {
					t.Fatalf("expected y %#v got %#v", tc.y, y)
				}
			}

			// Below we allocate slices to access items using the generated indices.
			// If we are able to access those items using any given index without
			// creating a runtime panic, then our implementation is safe to the extend
			// that no indices are generated out of bounds.
			{
				slx := make([]int, tc.a)
				for _, i := range x {
					_ = slx[i]
				}

				sly := make([]int, tc.b)
				for _, i := range y {
					_ = sly[i]
				}
			}
		})
	}
}

func musStr(lis []int) string {
	byt, err := json.Marshal(lis)
	if err != nil {
		tracer.Panic(tracer.Mask(err))
	}

	return string(byt)
}
