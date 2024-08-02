package postvalidator

import "github.com/uvio-network/apiserver/pkg/storage/poststorage"

type Interface interface {
	Create([]*poststorage.Object) ([]*poststorage.Object, error)
}
