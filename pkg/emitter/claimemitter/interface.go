package claimemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
)

type Interface interface {
	Create(objectid.ID, objectlabel.DesiredLifecycle) error
}
