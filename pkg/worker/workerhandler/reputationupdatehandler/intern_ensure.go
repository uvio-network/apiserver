package reputationupdatehandler

import (
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
)

func (h *InternHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	// TODO only calculate onchain reputation for claims using contract version >= v0.5.0
	//
	//     market                                   compute units
	//
	//     searchResults(true) (bool, ...)          10
	//     searchIndices(true) (uint256, ...)       10
	//
	//     competence                               compute units
	//
	//     searchStakers(true) []Address            10
	//     searchStakers(false) []Address           10
	//     searchHistory(true) []uint256            10
	//     searchHistory(false) []uint256           10
	//
	//     integrity                                compute units
	//
	//     searchVoters(true) []Address             10
	//     searchVoters(false) []Address            10
	//     searchSamples(true) []uint256            10
	//     searchSamples(false) []uint256           10
	//
	//     total                                    100
	//
	return nil
}
