package generic

import (
	"fmt"
	"slices"
	"testing"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

func Test_Generic_Unique_string(t *testing.T) {
	testCases := []struct {
		lis []string
		uni []string
	}{
		// Case 000
		{
			lis: []string{},
			uni: []string{},
		},
		// Case 001
		{
			lis: []string{
				"55",
				"44",
			},
			uni: []string{
				"44",
				"55",
			},
		},
		// Case 002
		{
			lis: []string{
				"33",
				"44",
				"33",
				"33",
			},
			uni: []string{
				"33",
				"44",
			},
		},
		// Case 003
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"33",
				"55",
				"66",
				"55",
				"88",
			},
			uni: []string{
				"22",
				"33",
				"44",
				"55",
				"66",
				"88",
			},
		},
		// Case 004
		{
			lis: []string{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
			uni: []string{
				"22",
				"33",
				"44",
				"55",
				"66",
				"88",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Unique(tc.lis)

			slices.Sort(uni)
			slices.Sort(tc.uni)

			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}

func Test_Generic_Unique_ID(t *testing.T) {
	testCases := []struct {
		lis []objectid.ID
		uni []objectid.ID
	}{
		// Case 000
		{
			lis: []objectid.ID{},
			uni: []objectid.ID{},
		},
		// Case 001
		{
			lis: []objectid.ID{
				"55",
				"44",
			},
			uni: []objectid.ID{
				"44",
				"55",
			},
		},
		// Case 002
		{
			lis: []objectid.ID{
				"33",
				"44",
				"33",
				"33",
			},
			uni: []objectid.ID{
				"33",
				"44",
			},
		},
		// Case 003
		{
			lis: []objectid.ID{
				"33",
				"44",
				"88",
				"22",
				"33",
				"55",
				"66",
				"55",
				"88",
			},
			uni: []objectid.ID{
				"22",
				"33",
				"44",
				"55",
				"66",
				"88",
			},
		},
		// Case 004
		{
			lis: []objectid.ID{
				"33",
				"44",
				"88",
				"22",
				"55",
				"66",
			},
			uni: []objectid.ID{
				"22",
				"33",
				"44",
				"55",
				"66",
				"88",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Unique(tc.lis)

			slices.Sort(uni)
			slices.Sort(tc.uni)

			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}

func Test_Generic_Unique_int64(t *testing.T) {
	testCases := []struct {
		lis []int64
		uni []int64
	}{
		// Case 000
		{
			lis: []int64{},
			uni: []int64{},
		},
		// Case 001
		{
			lis: []int64{
				55,
				44,
			},
			uni: []int64{
				44,
				55,
			},
		},
		// Case 002
		{
			lis: []int64{
				33,
				44,
				33,
				33,
			},
			uni: []int64{
				33,
				44,
			},
		},
		// Case 003
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				33,
				55,
				66,
				55,
				88,
			},
			uni: []int64{
				22,
				33,
				44,
				55,
				66,
				88,
			},
		},
		// Case 004
		{
			lis: []int64{
				33,
				44,
				88,
				22,
				55,
				66,
			},
			uni: []int64{
				22,
				33,
				44,
				55,
				66,
				88,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			uni := Unique(tc.lis)

			slices.Sort(uni)
			slices.Sort(tc.uni)

			if !slices.Equal(uni, tc.uni) {
				t.Fatalf("expected %#v got %#v", tc.uni, uni)
			}
		})
	}
}
