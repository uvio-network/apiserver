package userstorage

import (
	"fmt"
	"testing"
)

func Test_Storage_UserStorage_Object_Reputation(t *testing.T) {
	testCases := []struct {
		use *Object
		rep float64
	}{
		// Case 000
		{
			use: &Object{
				Summary: []float64{
					0, // Right
					0, // Wrong
					0, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 0,
		},
		// Case 001
		{
			use: &Object{
				Summary: []float64{
					0, // Right
					1, // Wrong
					1, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 0,
		},
		// Case 002
		{
			use: &Object{
				Summary: []float64{
					0, // Right
					1, // Wrong
					0, // Honest
					1, // Dishonest
					0, // Abstained
				},
			},
			rep: 0,
		},
		// Case 003
		{
			use: &Object{
				Summary: []float64{
					1, // Right
					0, // Wrong
					0, // Honest
					1, // Dishonest
					0, // Abstained
				},
			},
			rep: 0,
		},
		// Case 004
		{
			use: &Object{
				Summary: []float64{
					1, // Right
					0, // Wrong
					0, // Honest
					0, // Dishonest
					1, // Abstained
				},
			},
			rep: 0,
		},
		// Case 005
		{
			use: &Object{
				Summary: []float64{
					1, // Right
					0, // Wrong
					1, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 2,
		},
		// Case 006
		{
			use: &Object{
				Summary: []float64{
					4, // Right
					0, // Wrong
					2, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 3,
		},
		// Case 007
		{
			use: &Object{
				Summary: []float64{
					4, // Right
					1, // Wrong
					2, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 2.4,
		},
		// Case 008
		{
			use: &Object{
				Summary: []float64{
					4, // Right
					0, // Wrong
					2, // Honest
					0, // Dishonest
					1, // Abstained
				},
			},
			rep: 2.3,
		},
		// Case 009
		{
			use: &Object{
				Summary: []float64{
					4, // Right
					3, // Wrong
					2, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 1.71429,
		},
		// Case 010
		{
			use: &Object{
				Summary: []float64{
					4, // Right
					0, // Wrong
					2, // Honest
					0, // Dishonest
					3, // Abstained
				},
			},
			rep: 1.07143,
		},
		// Case 011
		{
			use: &Object{
				Summary: []float64{
					4, // Right
					1, // Wrong
					2, // Honest
					1, // Dishonest
					0, // Abstained
				},
			},
			rep: 1.33333,
		},
		// Case 012
		{
			use: &Object{
				Summary: []float64{
					1, // Right
					4, // Wrong
					2, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 0.6,
		},
		// Case 013
		{
			use: &Object{
				Summary: []float64{
					2, // Right
					2, // Wrong
					4, // Honest
					0, // Dishonest
					0, // Abstained
				},
			},
			rep: 2.5,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			rep := tc.use.Reputation()

			if rep != tc.rep {
				t.Fatalf("expected %#v got %#v", tc.rep, rep)
			}
		})
	}
}
