package validator

import (
	"github.com/uvio-network/apiserver/pkg/validator/postvalidator"
	"github.com/uvio-network/apiserver/pkg/validator/votevalidator"
)

// Interface provides cross resource validation for storage specific
// implementations. For instance the vote storage may need to use the post
// storage in order to verify the integrity of vote objects. At the same time
// the post storage may need to use the vote storage in order to verify the
// integrity of the post objects. The abstraction of the validator interface
// allows the validation of business logic that is cross dependent on multiple
// resources in either direction.
type Interface interface {
	Post() postvalidator.Interface
	Vote() votevalidator.Interface
}
