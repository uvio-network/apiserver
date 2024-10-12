package generic

import (
	"fmt"
	"slices"
	"testing"

	"github.com/xh3b4sd/objectid"
)

func Test_Generic_Difference_string(t *testing.T) {
	testCases := []struct {
		one []string
		two []string
		lis []string
	}{
		// Case 000
		{
			one: []string{},
			two: []string{},
			lis: []string{},
		},
		// Case 001
		{
			one: []string{},
			two: []string{
				"55",
				"44",
			},
			lis: []string{
				"55",
				"44",
			},
		},
		// Case 002
		{
			one: []string{
				"55",
				"44",
			},
			two: []string{},
			lis: []string{
				"55",
				"44",
			},
		},
		// Case 003
		{
			one: []string{
				"55",
				"44",
			},
			two: []string{
				"55",
				"44",
			},
			lis: []string{},
		},
		// Case 004
		{
			one: []string{
				"55",
				"44",
				"55",
			},
			two: []string{
				"44",
				"55",
			},
			lis: []string{},
		},
		// Case 005
		{
			one: []string{
				"11",
				"22",
				"33",
				"44",
				"33",
				"33",
			},
			two: []string{
				"33",
				"44",
				"33",
				"55",
			},
			lis: []string{
				"11",
				"22",
				"55",
			},
		},
		// Case 006
		{
			one: []string{
				"33",
				"44",
				"33",
				"55",
			},
			two: []string{
				"11",
				"22",
				"33",
				"44",
				"33",
				"33",
			},
			lis: []string{
				"11",
				"22",
				"55",
			},
		},
		// Case 007
		{
			one: []string{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			two: []string{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			lis: []string{
				"22",
				"66",
			},
		},
		// Case 008
		{
			one: []string{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			two: []string{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			lis: []string{
				"22",
				"66",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lis := Difference(tc.one, tc.two)

			slices.Sort(lis)
			slices.Sort(tc.lis)

			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}

func Test_Generic_Difference_objectID(t *testing.T) {
	testCases := []struct {
		one []objectid.ID
		two []objectid.ID
		lis []objectid.ID
	}{
		// Case 000
		{
			one: []objectid.ID{},
			two: []objectid.ID{},
			lis: []objectid.ID{},
		},
		// Case 001
		{
			one: []objectid.ID{},
			two: []objectid.ID{
				"55",
				"44",
			},
			lis: []objectid.ID{
				"55",
				"44",
			},
		},
		// Case 002
		{
			one: []objectid.ID{
				"55",
				"44",
			},
			two: []objectid.ID{},
			lis: []objectid.ID{
				"55",
				"44",
			},
		},
		// Case 003
		{
			one: []objectid.ID{
				"55",
				"44",
			},
			two: []objectid.ID{
				"55",
				"44",
			},
			lis: []objectid.ID{},
		},
		// Case 004
		{
			one: []objectid.ID{
				"55",
				"44",
				"55",
			},
			two: []objectid.ID{
				"44",
				"55",
			},
			lis: []objectid.ID{},
		},
		// Case 005
		{
			one: []objectid.ID{
				"11",
				"22",
				"33",
				"44",
				"33",
				"33",
			},
			two: []objectid.ID{
				"33",
				"44",
				"33",
				"55",
			},
			lis: []objectid.ID{
				"11",
				"22",
				"55",
			},
		},
		// Case 006
		{
			one: []objectid.ID{
				"33",
				"44",
				"33",
				"55",
			},
			two: []objectid.ID{
				"11",
				"22",
				"33",
				"44",
				"33",
				"33",
			},
			lis: []objectid.ID{
				"11",
				"22",
				"55",
			},
		},
		// Case 007
		{
			one: []objectid.ID{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			two: []objectid.ID{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			lis: []objectid.ID{
				"22",
				"66",
			},
		},
		// Case 008
		{
			one: []objectid.ID{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			two: []objectid.ID{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			lis: []objectid.ID{
				"22",
				"66",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lis := Difference(tc.one, tc.two)

			slices.Sort(lis)
			slices.Sort(tc.lis)

			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}

func Test_Generic_Difference_int64(t *testing.T) {
	testCases := []struct {
		one []int64
		two []int64
		lis []int64
	}{
		// Case 000
		{
			one: []int64{},
			two: []int64{},
			lis: []int64{},
		},
		// Case 001
		{
			one: []int64{},
			two: []int64{
				55,
				44,
			},
			lis: []int64{
				55,
				44,
			},
		},
		// Case 002
		{
			one: []int64{
				55,
				44,
			},
			two: []int64{},
			lis: []int64{
				55,
				44,
			},
		},
		// Case 003
		{
			one: []int64{
				55,
				44,
			},
			two: []int64{
				55,
				44,
			},
			lis: []int64{},
		},
		// Case 004
		{
			one: []int64{
				55,
				44,
				55,
			},
			two: []int64{
				44,
				55,
			},
			lis: []int64{},
		},
		// Case 005
		{
			one: []int64{
				11,
				22,
				33,
				44,
				33,
				33,
			},
			two: []int64{
				33,
				44,
				33,
				55,
			},
			lis: []int64{
				11,
				22,
				55,
			},
		},
		// Case 006
		{
			one: []int64{
				33,
				44,
				33,
				55,
			},
			two: []int64{
				11,
				22,
				33,
				44,
				33,
				33,
			},
			lis: []int64{
				11,
				22,
				55,
			},
		},
		// Case 007
		{
			one: []int64{
				22,
				44,
				11,
				33,
				55,
			},
			two: []int64{
				33,
				44,
				11,
				55,
				66,
				33,
			},
			lis: []int64{
				22,
				66,
			},
		},
		// Case 008
		{
			one: []int64{
				33,
				44,
				11,
				55,
				66,
				33,
			},
			two: []int64{
				22,
				44,
				11,
				33,
				55,
			},
			lis: []int64{
				22,
				66,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lis := Difference(tc.one, tc.two)

			slices.Sort(lis)
			slices.Sort(tc.lis)

			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}
