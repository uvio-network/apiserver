package objectfield

import (
	"fmt"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Lifecycle struct {
	// Data is the string data of this object field.
	Data objectlabel.DesiredLifecycle `json:"data,omitempty"`
	// Hash is the lifecycle confirmation of this object field.
	Hash []string `json:"hash,omitempty"`
	// Time is the most recent time at which this object field got updated.
	Time time.Time `json:"time,omitempty"`
}

func (l Lifecycle) Empty() bool {
	return l.Data == "" && len(l.Hash) == 0
}

func (l Lifecycle) Is(lis ...objectlabel.DesiredLifecycle) bool {
	for _, x := range lis {
		if l.Data == x {
			return true
		}
	}

	return false
}

func (l Lifecycle) Pending() bool {
	return len(l.Hash) == 0
}

// String returns the underlying lifecycle, formatted using its desired
// lifecycle phase and its current system status, separated by a colon, e.g.
// propose:pending or resolve:onchain.
func (l Lifecycle) String() string {
	if l.Empty() {
		return ""
	}

	if len(l.Hash) == 0 {
		if l.Data == objectlabel.LifecycleOnchain {
			return string(objectlabel.LifecyclePending)
		}

		return fmt.Sprintf("%s:%s", l.Data, objectlabel.LifecyclePending)
	}

	if l.Data == objectlabel.LifecycleOnchain {
		return string(objectlabel.LifecycleOnchain)
	}

	return fmt.Sprintf("%s:%s", l.Data, objectlabel.LifecycleOnchain)
}
