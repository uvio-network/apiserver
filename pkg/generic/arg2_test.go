package generic

import (
	"fmt"
	"slices"
	"testing"

	"github.com/xh3b4sd/objectid"
)

func Test_ObjectID_Arg2_string(t *testing.T) {
	testCases := []struct {
		one []string
		two []string
		str string
		key []string
	}{
		// Case 000
		{
			one: []string{
				"foo",
				"bar",
			},
			two: []string{
				"baz",
				"zap",
			},
			str: "des/eve/%s/lis/%s",
			key: []string{
				"des/eve/foo/lis/baz",
				"des/eve/bar/lis/zap",
			},
		},
		// Case 001
		{
			one: []string{
				"foo",
				"bar",
			},
			two: []string{
				"baz",
				"zap",
			},
			str: fmt.Sprintf("wal/use/%s/obj/%s/lis/%s", "1234", "%s", "%s"),
			key: []string{
				"wal/use/1234/obj/foo/lis/baz",
				"wal/use/1234/obj/bar/lis/zap",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Arg2(tc.str, tc.one, tc.two)
			if !slices.Equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}

func Test_ObjectID_Arg2_ID(t *testing.T) {
	testCases := []struct {
		one []objectid.ID
		two []objectid.ID
		str string
		key []string
	}{
		// Case 000
		{
			one: []objectid.ID{
				"foo",
				"bar",
			},
			two: []objectid.ID{
				"baz",
				"zap",
			},
			str: "des/eve/%s/lis/%s",
			key: []string{
				"des/eve/foo/lis/baz",
				"des/eve/bar/lis/zap",
			},
		},
		// Case 001
		{
			one: []objectid.ID{
				"foo",
				"bar",
			},
			two: []objectid.ID{
				"baz",
				"zap",
			},
			str: fmt.Sprintf("wal/use/%s/obj/%s/lis/%s", "1234", "%s", "%s"),
			key: []string{
				"wal/use/1234/obj/foo/lis/baz",
				"wal/use/1234/obj/bar/lis/zap",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Arg2(tc.str, tc.one, tc.two)
			if !slices.Equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}
