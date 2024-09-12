package uvxminthandler

import (
	"fmt"
	"testing"
)

func Test_Worker_WorkerHandler_UVXMintHandler_InternEnsure_createPaging(t *testing.T) {
	testCases := []struct {
		cur string
		des string
	}{
		// Case 000
		{
			cur: "",
			des: "1",
		},
		// Case 001
		{
			cur: "0",
			des: "1",
		},
		// Case 002
		{
			cur: " ",
			des: "1",
		},
		// Case 003
		{
			cur: "foo",
			des: "1",
		},
		// Case 004
		{
			cur: "true",
			des: "1",
		},
		// Case 005
		{
			cur: "923 f32u d2",
			des: "1",
		},
		// Case 006
		{
			cur: "1",
			des: "2",
		},
		// Case 007
		{
			cur: "2",
			des: "0",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			des := createPaging(tc.cur)

			if des != tc.des {
				t.Fatalf("expected %#v got %#v", tc.des, des)
			}
		})
	}
}
