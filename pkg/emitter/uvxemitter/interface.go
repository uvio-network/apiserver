package uvxemitter

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	Mint(use objectid.ID) error
}
