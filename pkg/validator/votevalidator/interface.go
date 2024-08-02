package votevalidator

import "github.com/uvio-network/apiserver/pkg/storage/votestorage"

type Interface interface {
	Create([]*votestorage.Object) ([]*votestorage.Object, error)
}
