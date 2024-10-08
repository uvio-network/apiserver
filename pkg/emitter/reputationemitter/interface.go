package reputationemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

type Interface interface {
	// CompetenceUpdate emits an event that intends to trigger the update of the
	// competence specific user metrics for all users that participated in the
	// given market.
	//
	//     inp[0] the claim ID of the settled balance
	//
	CompetenceUpdate(objectid.ID) error

	// IntegritiyUpdate emits an event that intends to trigger the update of the
	// integritiy specific user metrics for all users that participated in the
	// given market.
	//
	//     inp[0] the claim ID of the settled balance
	//
	IntegrityUpdate(objectid.ID) error
}
