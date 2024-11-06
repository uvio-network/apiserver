package storage

import (
	"github.com/uvio-network/apiserver/pkg/storage/notestorage"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/userstorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
	"github.com/uvio-network/apiserver/pkg/storage/walletstorage"
)

type Interface interface {
	Note() notestorage.Interface
	Post() poststorage.Interface
	User() userstorage.Interface
	Vote() votestorage.Interface
	Wallet() walletstorage.Interface
}
