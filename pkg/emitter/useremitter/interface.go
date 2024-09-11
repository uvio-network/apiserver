package useremitter

import "github.com/uvio-network/apiserver/pkg/object/objectid"

type Interface interface {
	Create(use objectid.ID) error
}
