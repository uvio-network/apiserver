package reconciler

import (
	"github.com/uvio-network/apiserver/pkg/reconciler/postreconciler"
	"github.com/uvio-network/apiserver/pkg/reconciler/votereconciler"
	"github.com/uvio-network/apiserver/pkg/reconciler/walletreconciler"
)

// Interface provides cross resource reconciliation to bridge between server and
// storage specific implementations. For instance the vote storage may need to
// use the post storage in order to verify the integrity of vote objects. At the
// same time the post storage may need to use the vote storage in order to
// verify the integrity of the post objects. The abstraction of the reconciler
// interface allows for business logic that is cross dependent on multiple
// resources in either direction.
type Interface interface {
	Post() postreconciler.Interface
	Vote() votereconciler.Interface
	Wallet() walletreconciler.Interface
}
