package contract

import "github.com/uvio-network/apiserver/pkg/contract/uvxcontract"

type Interface interface {
	UVX() uvxcontract.Interface
}
