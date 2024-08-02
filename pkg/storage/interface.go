package storage

import (
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
)

type Interface interface {
	Post() poststorage.Interface
	User() userstorage.Interface
	Vote() votestorage.Interface
}
