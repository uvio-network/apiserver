package generic

import (
	"fmt"
	"slices"
	"testing"

	"github.com/xh3b4sd/objectid"
)

func Test_ObjectID_Arg1_string(t *testing.T) {
	testCases := []struct {
		one []string
		str string
		key []string
	}{
		// Case 000
		{
			one: []string{
				"foo",
				"bar",
			},
			str: "des/eve/%s",
			key: []string{
				"des/eve/foo",
				"des/eve/bar",
			},
		},
		// Case 001
		{
			one: []string{
				"foo",
				"bar",
			},
			str: fmt.Sprintf("wal/use/%s/obj/%s", "1234", "%s"),
			key: []string{
				"wal/use/1234/obj/foo",
				"wal/use/1234/obj/bar",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Arg1(tc.str, tc.one)
			if !slices.Equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}

func Test_ObjectID_Arg1_ID(t *testing.T) {
	testCases := []struct {
		one []objectid.ID
		str string
		key []string
	}{
		// Case 000
		{
			one: []objectid.ID{
				"foo",
				"bar",
			},
			str: "des/eve/%s",
			key: []string{
				"des/eve/foo",
				"des/eve/bar",
			},
		},
		// Case 001
		{
			one: []objectid.ID{
				"foo",
				"bar",
			},
			str: fmt.Sprintf("wal/use/%s/obj/%s", "1234", "%s"),
			key: []string{
				"wal/use/1234/obj/foo",
				"wal/use/1234/obj/bar",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			key := Arg1(tc.str, tc.one)
			if !slices.Equal(key, tc.key) {
				t.Fatalf("expected %#v got %#v", tc.key, key)
			}
		})
	}
}
