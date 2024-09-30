package hexencoding

import (
	"fmt"
	"testing"
)

func Test_Format_HexEncoding_Verify_failure(t *testing.T) {
	testCases := []struct {
		lis []string
	}{
		// Case 000
		{
			lis: []string{
				"",
				" ",
				"   0xC76C054AB91220Cd8BE3f96f4dFCDD4E40075C89     ",
			},
		},
		// Case 001
		{
			lis: []string{
				"foo",
				"bar!",
				"crapto",
			},
		},
		// Case 002
		{
			lis: []string{
				"0",
				"0x",
				"xdceb358e44d0e9a5a3d6f235ea79749bbe7d732201b25e082c1e3d5924ca6e84",
				"b358e44d0e9a5a3d6f235ea79749bbe7d732201b25e082c1e3d5924ca6e84",
				"0 x dceb358e44d0e9a5a3d6f235ea79749bbe7d732201b25e082c1e3d5924ca6e84",
			},
		},
		// Case 003
		{
			lis: []string{
				"0x1234M",
				"0 x 0857d6Cdb4aD4674fE9E7C9a5e9662D4Ad735Caa",
				"0xdceb358e44d0e9a5a3d6f 235ea79749bbe7d732201b25e082c1e3d5 924ca6e84",
				"0x0437c4df64cdef106fe01c0c55a579d05a78bb97fc4151840ed712f15 4407a01e07c91b07da6d1bf5ffa4930b941f4787b44c2c7b88e1efd8da2905df5cbd59cda",
			},
		},
		// Case 004
		{
			lis: []string{
				"0x0000000000000000000000000000000000000000",
				"0x0000000000000000000000000000000000000000000000000000000000000000",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			for _, x := range tc.lis {
				ver := Verify(x)
				if ver != false {
					t.Fatalf("expected %#v got %#v", false, ver)
				}
			}
		})
	}
}

func Test_Format_HexEncoding_Verify_success(t *testing.T) {
	testCases := []struct {
		lis []string
	}{
		// Case 000
		{
			lis: []string{
				"0x0857d6Cdb4aD4674fE9E7C9a5e9662D4Ad735Caa",
				"0x5C10b0E8920EC7F2c342B68025Aa3A86aF4C9b0c",
				"0xC76C054AB91220Cd8BE3f96f4dFCDD4E40075C89",
			},
		},
		// Case 001
		{
			lis: []string{
				"0xe8b2a9781676Fd22E9e2553dc4C04E134e2d3A96",
				"0x0000aEaD33B0dC6af586e9fCfdd503171dd81Ca7",
				"0x0000d35621A9172f07aC09301C5B6a9deAE60000",
				"0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			},
		},
		// Case 002
		{
			lis: []string{
				"0x0857d6Cdb4aD4674fE9E7C9a5e9662D4Ad735Caa",
				"0xdceb358e44d0e9a5a3d6f235ea79749bbe7d732201b25e082c1e3d5924ca6e84",
				"0x0437c4df64cdef106fe01c0c55a579d05a78bb97fc4151840ed712f154407a01e07c91b07da6d1bf5ffa4930b941f4787b44c2c7b88e1efd8da2905df5cbd59cda",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			for _, x := range tc.lis {
				ver := Verify(x)
				if ver != true {
					t.Fatalf("expected %#v got %#v", false, ver)
				}
			}
		})
	}
}
