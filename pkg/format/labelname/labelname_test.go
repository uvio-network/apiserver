package labelname

import (
	"fmt"
	"testing"
)

func Test_Format_LabelName_Format(t *testing.T) {
	testCases := []struct {
		str string
		trm string
	}{
		// Case 000
		{
			str: "",
			trm: "",
		},
		// Case 001
		{
			str: "foo",
			trm: "foo",
		},
		// Case 002
		{
			str: "hello world",
			trm: "hello-world",
		},
		// Case 003
		{
			str: " hello    world  ",
			trm: "hello-world",
		},
		// Case 004
		{
			str: " \t\n Hello     Gophers \n\t\r\n",
			trm: "hello-gophers",
		},
		// Case 005
		{
			str: " 030\t\naka     030 \n\t\r\n",
			trm: "030-aka-030",
		},
		// Case 006
		{
			str: "MEV",
			trm: "mev",
		},
		// Case 007
		{
			str: " MEV",
			trm: "mev",
		},
		// Case 008
		{
			str: "DeFi",
			trm: "defi",
		},
		// Case 009
		{
			str: "DeFi ",
			trm: "defi",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			trm := Format(tc.str)
			if trm != tc.trm {
				t.Fatalf("expected %#v got %#v", tc.trm, trm)
			}
		})
	}
}
