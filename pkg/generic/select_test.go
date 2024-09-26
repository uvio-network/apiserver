package generic

import (
	"fmt"
	"slices"
	"testing"

	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

func Test_Generic_Select_string(t *testing.T) {
	testCases := []struct {
		all []string
		sub []string
		lis []string
	}{
		// Case 000
		{
			all: []string{},
			sub: []string{},
			lis: []string{},
		},
		// Case 001
		{
			all: []string{
				"55",
				"44",
			},
			sub: []string{},
			lis: []string{},
		},
		// Case 002
		{
			all: []string{},
			sub: []string{
				"55",
				"44",
			},
			lis: []string{
				"44",
				"55",
			},
		},
		// Case 003
		{
			all: []string{
				"55",
				"44",
			},
			sub: []string{
				"55",
				"44",
			},
			lis: []string{},
		},
		// Case 004
		{
			all: []string{
				"44",
				"55",
			},
			sub: []string{
				"55",
				"44",
				"55",
			},
			lis: []string{},
		},
		// Case 005
		{
			all: []string{
				"33",
				"44",
				"33",
				"55",
			},
			sub: []string{
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
			},
		},
		// Case 006
		{
			all: []string{
				"11",
				"22",
				"33",
				"44",
				"33",
				"33",
			},
			sub: []string{
				"33",
				"44",
				"33",
				"55",
			},
			lis: []string{
				"55",
			},
		},
		// Case 007
		{
			all: []string{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			sub: []string{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			lis: []string{
				"22",
			},
		},
		// Case 008
		{
			all: []string{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			sub: []string{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			lis: []string{
				"66",
			},
		},
		// Case 009
		{
			all: []string{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			sub: []string{
				"66",
				"44",
				"11",
				"11",
				"66",
				"66",
			},
			lis: []string{
				"66",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lis := Select(tc.all, tc.sub)

			slices.Sort(lis)
			slices.Sort(tc.lis)

			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}

func Test_Generic_Select_objectID(t *testing.T) {
	testCases := []struct {
		all []objectid.ID
		sub []objectid.ID
		lis []objectid.ID
	}{
		// Case 000
		{
			all: []objectid.ID{},
			sub: []objectid.ID{},
			lis: []objectid.ID{},
		},
		// Case 001
		{
			all: []objectid.ID{
				"55",
				"44",
			},
			sub: []objectid.ID{},
			lis: []objectid.ID{},
		},
		// Case 002
		{
			all: []objectid.ID{},
			sub: []objectid.ID{
				"55",
				"44",
			},
			lis: []objectid.ID{
				"44",
				"55",
			},
		},
		// Case 003
		{
			all: []objectid.ID{
				"55",
				"44",
			},
			sub: []objectid.ID{
				"55",
				"44",
			},
			lis: []objectid.ID{},
		},
		// Case 004
		{
			all: []objectid.ID{
				"44",
				"55",
			},
			sub: []objectid.ID{
				"55",
				"44",
				"55",
			},
			lis: []objectid.ID{},
		},
		// Case 005
		{
			all: []objectid.ID{
				"33",
				"44",
				"33",
				"55",
			},
			sub: []objectid.ID{
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
			},
		},
		// Case 006
		{
			all: []objectid.ID{
				"11",
				"22",
				"33",
				"44",
				"33",
				"33",
			},
			sub: []objectid.ID{
				"33",
				"44",
				"33",
				"55",
			},
			lis: []objectid.ID{
				"55",
			},
		},
		// Case 007
		{
			all: []objectid.ID{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			sub: []objectid.ID{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			lis: []objectid.ID{
				"22",
			},
		},
		// Case 008
		{
			all: []objectid.ID{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			sub: []objectid.ID{
				"33",
				"44",
				"11",
				"55",
				"66",
				"33",
			},
			lis: []objectid.ID{
				"66",
			},
		},
		// Case 009
		{
			all: []objectid.ID{
				"22",
				"44",
				"11",
				"33",
				"55",
			},
			sub: []objectid.ID{
				"66",
				"44",
				"11",
				"11",
				"66",
				"66",
			},
			lis: []objectid.ID{
				"66",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lis := Select(tc.all, tc.sub)

			slices.Sort(lis)
			slices.Sort(tc.lis)

			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}

func Test_Generic_Select_int64(t *testing.T) {
	testCases := []struct {
		all []int64
		sub []int64
		lis []int64
	}{
		// Case 000
		{
			all: []int64{},
			sub: []int64{},
			lis: []int64{},
		},
		// Case 001
		{
			all: []int64{
				55,
				44,
			},
			sub: []int64{},
			lis: []int64{},
		},
		// Case 002
		{
			all: []int64{},
			sub: []int64{
				55,
				44,
			},
			lis: []int64{
				44,
				55,
			},
		},
		// Case 003
		{
			all: []int64{
				55,
				44,
			},
			sub: []int64{
				55,
				44,
			},
			lis: []int64{},
		},
		// Case 004
		{
			all: []int64{
				44,
				55,
			},
			sub: []int64{
				55,
				44,
				55,
			},
			lis: []int64{},
		},
		// Case 005
		{
			all: []int64{
				33,
				44,
				33,
				55,
			},
			sub: []int64{
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
			},
		},
		// Case 006
		{
			all: []int64{
				11,
				22,
				33,
				44,
				33,
				33,
			},
			sub: []int64{
				33,
				44,
				33,
				55,
			},
			lis: []int64{
				55,
			},
		},
		// Case 007
		{
			all: []int64{
				33,
				44,
				11,
				55,
				66,
				33,
			},
			sub: []int64{
				22,
				44,
				11,
				33,
				55,
			},
			lis: []int64{
				22,
			},
		},
		// Case 008
		{
			all: []int64{
				22,
				44,
				11,
				33,
				55,
			},
			sub: []int64{
				33,
				44,
				11,
				55,
				66,
				33,
			},
			lis: []int64{
				66,
			},
		},
		// Case 009
		{
			all: []int64{
				22,
				44,
				11,
				33,
				55,
			},
			sub: []int64{
				66,
				44,
				11,
				11,
				66,
				66,
			},
			lis: []int64{
				66,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lis := Select(tc.all, tc.sub)

			slices.Sort(lis)
			slices.Sort(tc.lis)

			if !slices.Equal(lis, tc.lis) {
				t.Fatalf("expected %#v got %#v", tc.lis, lis)
			}
		})
	}
}
