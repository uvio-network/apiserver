package reputationemitter

import (
	"github.com/uvio-network/apiserver/pkg/object/objectid"
)

type Interface interface {
	// CompetenceUpdate emits an event that intends to trigger the update of the
	// competence specific user metrics for all users that participated in the
	// given market.
	//
	//     inp[0] the block number at which the call to CompetenceUpdate was made
	//     inp[1] the claim ID of the settled balance
	//
	CompetenceUpdate(uint64, objectid.ID) error

	// IntegritiyUpdate emits an event that intends to trigger the update of the
	// integritiy specific user metrics for all users that participated in the
	// given market.
	//
	//     inp[0] the block number at which the call to IntegrityUpdate was made
	//     inp[1] the claim ID of the settled balance
	//
	IntegrityUpdate(uint64, objectid.ID) error
}
