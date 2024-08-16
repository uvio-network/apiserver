package claimresolvehandler

import (
	"fmt"
	"math/big"
	"strconv"

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

	//     find claims with lifecycle "propose" that are expired
	var cla []*poststorage.Object
	{
		cla, err = h.sto.Post().SearchExpiry()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	for _, x := range cla {
		var tre *big.Int
		{
			tre, err = metToTre(x.Meta)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		var det [4]marketscontract.IMarketsClaim
		{
			det, err = h.mar.Claims(nil, tre)
			if err != nil {
				return tracer.Mask(err)
			}
		}

		fmt.Printf("det %#v\n", det)

		// read from chain if claim has to be resolved, just double checking

		// 1. claim is not yet resolvable, this is a bug, we just log an error,
		// maybe send an email

		// 2. claim was already resolved, only remove claim from Redis

		// last step when everything is nice and clean, remove the claim ID from
		// Redis post/expiry
	}

	return nil
}

func metToTre(met string) (*big.Int, error) {
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
