package claimscontract

import "github.com/ethereum/go-ethereum/common"

var contract = map[int]map[string]string{
	1337: {
		"0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0": "v0.4.0",
		"0x0165878A594ca255338adfa4d48449f69242Eb8F": "v0.5.0",
	},
	84532: {
		"0x537cE8e9F4Cce5a1D8033B63f274187157a966b3": "v0.4.0",
		"0x206ab72edea55819a9a90622873976A79d3419E3": "v0.5.0",
	},
}

type Contract struct {
	Address common.Address
	Version string
}

func NewContract(cid int, add string) *Contract {
	for k, v := range contract[cid] {
		if k == add {
			return &Contract{
				Address: common.HexToAddress(k),
				Version: v,
			}
		}
	}

	return nil
}
