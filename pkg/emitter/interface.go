package emitter

import "github.com/uvio-network/apiserver/pkg/emitter/useremitter"

type Interface interface {
	User() useremitter.Interface
}
