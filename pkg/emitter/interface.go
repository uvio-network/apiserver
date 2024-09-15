package emitter

import (
	"github.com/uvio-network/apiserver/pkg/emitter/claimemitter"
	"github.com/uvio-network/apiserver/pkg/emitter/useremitter"
	"github.com/uvio-network/apiserver/pkg/emitter/uvxemitter"
)

type Interface interface {
	Claim() claimemitter.Interface
	User() useremitter.Interface
	UVX() uvxemitter.Interface
}
