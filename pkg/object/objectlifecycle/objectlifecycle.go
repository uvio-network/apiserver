package objectlifecycle

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Lifecycle struct {
	// Data is the string data of this object field.
	Data string `json:"data"`
	// Hash is the lifecycle confirmation of this object field.
	Hash string `json:"hash"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time"`
}

func (l Lifecycle) Empty() bool {
	return l.Data == "" && l.Hash == ""
}

func (l Lifecycle) Is(lis ...string) bool {
	for _, x := range lis {
		if l.Data == x {
			return true
		}
	}

	return false
}

func (l Lifecycle) String() string {
	if l.Hash == "" {
		return objectlabel.LifecyclePending
	}

	return l.Data
}
