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
			rep: 1,
		},
		// Case 006
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
			rep: 0.8,
		},
		// Case 007 shows that abstaining is as detrimental to reputation as is
		// being wrong.
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
			rep: 0.8,
		},
		// Case 008
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
			rep: 0.5714285714285714,
		},
		// Case 009 shows that abstaining is as detrimental to reputation as is
		// being wrong.
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
			rep: 0.5714285714285714,
		},
		// Case 010
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
			rep: 0.5333333333333333,
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
