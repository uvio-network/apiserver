package claimresolvehandler

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/uvio-network/apiserver/pkg/contract/marketscontract"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) error {
	var err error

	// find trees with lifecycle "propose" that are expired
	var trees []*poststorage.Object
	{
		trees, err = h.sto.Post().SearchExpiry()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range trees {
		var treeId *big.Int
		{
			treeId, err = metaToTreeId(x.Meta)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var claims [4]marketscontract.IMarketsClaim
		{
			claims, err = h.markets.Claims(nil, treeId)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var claimsLength *big.Int
		{
			claimsLength, err = h.markets.ClaimsLength(nil, treeId)
			if err != nil {
				return tracer.Mask(err)
			}

			// sanity check - we cap the number of claims to 4 in the contract
			if claimsLength.Cmp(big.NewInt(4)) == 1 {
				return tracer.Maskf(runtime.ExecutionFailedError, "claimresolverhandler: invalid state: too many claims")
			}
		}

		var claim marketscontract.IMarketsClaim
		{
			claim = claims[claimsLength.Int64()-1]
		}

		if claim.Status == 1 { // claim.Status == Active

			var yeaVoters []string
			var nayVoters []string
			yeaStakersLength := len(claim.Stake.YeaStakers)
			nayStakersLength := len(claim.Stake.NayStakers)
			if yeaStakersLength > 0 && nayStakersLength > 0 {
				rand.Seed(time.Now().UnixNano())
				votersLength := int(math.Min(float64(yeaStakersLength), float64(nayStakersLength)))
				for i := 0; i < votersLength; i++ {
					randomIndex := rand.Intn(yeaStakersLength)
					yeaVoters = append(yeaVoters, claim.Stake.YeaStakers[randomIndex].Hex())

					randomIndex = rand.Intn(nayStakersLength)
					nayVoters = append(nayVoters, claim.Stake.NayStakers[randomIndex].Hex())
				}
			} else if yeaStakersLength > 0 && nayStakersLength == 0 {
				randomIndex := rand.Intn(yeaStakersLength)
				yeaVoters = append(yeaVoters, claim.Stake.YeaStakers[randomIndex].Hex())
			} else if nayStakersLength > 0 && yeaStakersLength == 0 {
				randomIndex := rand.Intn(nayStakersLength)
				nayVoters = append(nayVoters, claim.Stake.NayStakers[randomIndex].Hex())
			} else {
				return tracer.Maskf(runtime.ExecutionFailedError, "claimresolverhandler: invalid state: invalid stakers")
			}

			fmt.Println("yeaVoters: ", yeaVoters)
			fmt.Println("nayVoters: ", nayVoters)

		} else if claim.Status == 2 { // claim.Status == PendingVote
			fmt.Println("claimresolverhandler: claim has already been resolved")
		} else { // this should never happen
			return tracer.Maskf(runtime.ExecutionFailedError, "claimresolverhandler: invalid state: invalid claim status")
		}

		// 2. claim was already resolved, only remove claim from Redis

		// last step when everything is nice and clean, remove the claim ID from
		// Redis post/expiry
	}

	return nil
}

func metaToTreeId(met string) (*big.Int, error) {
	var err error

	var spl []string
	{
		spl = converter.StringToSlice(met)
	}

	if len(spl) == 0 {
		return nil, tracer.Maskf(runtime.ExecutionFailedError, "tree ID not found in meta data")
	}

	// We take the first item here because the first item is the onchain tree ID
	// that clients put in the meta data when creating post objects of kind
	// "claim".
	//
	//     "meta": "9,0"
	//
	var tre int64
	{
		tre, err = strconv.ParseInt(spl[0], 10, 64)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	return big.NewInt(tre), nil
}