package postreconciler

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
)

type Interface interface {
	// CreateBalance creates a new pending balance for the provided resolve.
	//
	//     inp[0] the resolve to create a pending balance for
	//     out[0] the pending balance as persisted in the underlying storage
	//
	CreateBalance(*poststorage.Object) (*poststorage.Object, error)

	// CreatePost prepares the provided post objects so that they can be persisted
	// in the underlying storage the first time.
	//
	//     inp[0] the post objects according to the raw user input
	//     out[0] the post objects verified, extended and ready for storage
	//
	CreatePost([]*poststorage.Object) ([]*poststorage.Object, error)

	// CreateResolve creates a new pending resolve for the provided propose using
	// the provided expiry.
	//
	//     inp[0] the propose to create a pending resolve for
	//     inp[1] the desired resolve expiry
	//     out[0] the pending resolve as persisted in the underlying storage
	//
	CreateResolve(*poststorage.Object, time.Time) (*poststorage.Object, error)

	// DeletePost prepares the provided post objects so that they can be removed
	// from the underlying storage.
	//
	//     inp[0] the post objects to be deleted
	//     out[0] the post objects ready to be removed from storage
	//
	DeletePost(inp []*poststorage.Object) ([]*poststorage.Object, error)

	// UpdateBalance modifies the provided post object of lifecycle phase
	// "balance" with the given transaction hashes. Once UpdateBalance succeeded,
	// the provided balance is not considered pending anymore, since the provided
	// transaction hashes got confirmed onchain.
	//
	//     inp[0] the post object of lifecycle phase "balance"
	//     inp[1] the confirmed onchain transaction hashes to set
	//
	UpdateBalance(*poststorage.Object, []common.Hash) error

	// UpdateHash modifies the transaction hash of the claims as provided by the
	// given post objects. Note that transaction hashes can only be updated once,
	// if their prior value was empty.
	//
	//     inp[0] the post objects of the claims to modify
	//     inp[1] the transaction hashes to set as the new values
	//     out[0] the post objects reflecting their updated state
	//
	UpdateHash([]*poststorage.Object, []string) ([]*poststorage.Object, error)

	// UpdateMeta modifies the post meta of the provided post objects.
	//
	//     inp[0] the post objects of the claims to modify
	//     inp[1] the meta data to set as the new values
	//     out[0] the post objects reflecting their updated state
	//
	UpdateMeta([]*poststorage.Object, []string) ([]*poststorage.Object, error)

	// UpdateResolve modifies the provided post object of lifecycle phase
	// "resolve" with the given transaction hash and voter addresses. Once
	// UpdateResolve succeeded, the provided resolve is not considered pending
	// anymore, since the provided transaction hash got confirmed onchain.
	//
	//     inp[0] the post object of lifecycle phase "resolve"
	//     inp[1] the confirmed onchain transaction hash to set
	//     inp[2] the staker addresses to set
	//
	UpdateResolve(*poststorage.Object, common.Hash, []common.Address) error

	// UpdateVotes modifies the vote summary of the existing post objects that are
	// linked to the provided onchain vote objects, so that said vote summary
	// properly reflects the votes being cast. UpdateVotes may not return any post
	// objects if none of the provided vote objects are confirmed onchain yet.
	//
	//     inp[0] the vote objects of the votes being cast
	//     out[0] the post objects reflecting their updated vote summary
	//
	UpdateVotes([]*votestorage.Object) ([]*poststorage.Object, error)
}
