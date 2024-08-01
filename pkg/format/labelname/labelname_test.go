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

func Test_Format_LabelName_Verify_failure(t *testing.T) {
	testCases := []struct {
		lis []string
	}{
		// Case 000
		{
			lis: []string{
				"foo",
				"bar!",
				"crapto",
			},
		},
		// Case 001
		{
			lis: []string{
				".foo",
				"fr.Ance",
				"Cryp.to",
			},
		},
		// Case 002
		{
			lis: []string{
				"$hello world",
				" hello    world  ",
				"99%",
			},
		},
		// Case 003
		{
			lis: []string{
				"ballin' ",
			},
		},
		// Case 004
		{
			lis: []string{
				"0.5%",
				"!!!",
			},
		},
		// Case 005
		{
			lis: []string{
				"De-Fi.to",
				"(DeFi)",
			},
		},
		// Case 006
		{
			lis: []string{
				"looks like winning/breh",
				"Is this a question?",
			},
		},
		// Case 007
		{
			lis: []string{
				"nasty",
				"?",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ver := Verify(tc.lis)
			if ver != false {
				t.Fatalf("expected %#v got %#v", false, ver)
			}
		})
	}
}

func Test_Format_LabelName_Verify_success(t *testing.T) {
	testCases := []struct {
		lis []string
	}{
		// Case 000
		{
			lis: []string{},
		},
		// Case 001
		{
			lis: []string{
				"foo",
				"frAnce 007",
				"Crypto47",
			},
		},
		// Case 002
		{
			lis: []string{
				"hello world",
				"hello - world",
				" hello    world  ",
			},
		},
		// Case 003
		{
			lis: []string{
				" \t\n Hello     Gophers \n\t\r\n",
				" 030\t\naka     030 \n\t\r\n",
			},
		},
		// Case 004
		{
			lis: []string{
				"MEV",
				"M-E-V-95",
				" MEV",
			},
		},
		// Case 005
		{
			lis: []string{
				"DeFi-",
				"DeFi ",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ver := Verify(tc.lis)
			if ver != true {
				t.Fatalf("expected %#v got %#v", true, ver)
			}
		})
	}
}
