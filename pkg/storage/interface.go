package storage

import "github.com/uvio-network/apiserver/pkg/storage/poststorage"

type Interface interface {
	Post() poststorage.Interface
}
