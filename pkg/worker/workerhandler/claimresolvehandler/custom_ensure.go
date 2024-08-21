package claimresolvehandler

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/uvio-network/apiserver/pkg/contract/marketscontract"
	randomizer "github.com/uvio-network/apiserver/pkg/contract/randomizercontract"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/runtime"
	"github.com/uvio-network/apiserver/pkg/server/converter"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/worker/budget"
	"github.com/xh3b4sd/rescue/task"
	"github.com/xh3b4sd/tracer"
)

func (h *SystemHandler) Ensure(tas *task.Task, bud *budget.Budget) (err error) {
	// find claims with lifecycle "propose" that are expired
	claims, err := h.sto.Post().SearchExpiry(objectlabel.LifecyclePropose)
	if err != nil {
		return tracer.Mask(err)
	}

	for _, x := range claims {
		// get the corresponding claim object from chain
		claim, treeId, err := h.getClaim(x)
		if err != nil {
			return tracer.Mask(err)
		}

		// handle claim according to its onchain Status
		if claim.Status == 1 { // claim.Status == Active
			// choose random voters and write to chain
			err = h.prepareVote(claim, treeId)
			if err != nil {
				return tracer.Mask(err)
			}
		} else if claim.Status == 2 { // claim.Status == PendingVote
			h.log.Log(
				context.Background(),
				"level", "info",
				"message", "claimresolverhandler: claim has already been resolved",
			)
		} else { // this should never happen
			h.log.Log(
				context.Background(),
				"level", "info",
				"message", "claimresolverhandler: invalid state: invalid claim status",
			)
		}

		// remove the claim ID from Redis post/expiry
		err = h.sto.Post().DeleteExpiry([]*poststorage.Object{x})
		if err != nil {
			return tracer.Mask(err)
		}

		// @todo - here
		// create a new post object with Lifecycle resolve and the same tree id
		// inp = append(inp, &poststorage.Object{
		// 	Chain:  x.Public.Chain,
		// 	Expiry: converter.StringToTime(x.Public.Expiry),
		// 	Kind:   x.Public.Kind,
		// 	Labels: converter.StringToSlice(x.Public.Labels),
		// 	Lifecycle: objectfield.Lifecycle{
		// 		Data: objectlabel.DesiredLifecycle(x.Public.Lifecycle),
		// 		Hash: x.Public.Hash,
		// 	},
		// 	Meta:   x.Public.Meta,
		// 	Owner:  userid.FromContext(ctx),
		// 	Parent: objectid.ID(x.Public.Parent),
		// 	Text:   x.Public.Text,
		// 	Token:  x.Public.Token,
		// })

		// Chain     string                `json:"chain"`
		// Created   time.Time             `json:"created"`
		// Expiry    time.Time             `json:"expiry"`
		// ID        objectid.ID           `json:"id"`
		// Kind      string                `json:"kind"`
		// Labels    []string              `json:"labels"`
		// Lifecycle objectfield.Lifecycle `json:"lifecycle"`
		// Meta      string                `json:"meta"`
		// Owner     objectid.ID           `json:"owner"`
		// Parent    objectid.ID           `json:"parent"`
		// Text      string                `json:"text"`
		// Token     string                `json:"token"`
		// Tree      objectid.ID           `json:"tree"`
		// Votes     []float64             `json:"votes"`
		// newPost :=
		// err = h.sto.Post().CreatePost([]*poststorage.Object{x})
		// if err != nil {
		// 	return tracer.Mask(err)
		// }
	}

	return nil
}

func (h *SystemHandler) getClaim(x *poststorage.Object) (claim marketscontract.IMarketsClaim, treeId *big.Int, err error) {
	treeId, err = metaToTreeId(x.Meta)
	if err != nil {
		return claim, treeId, tracer.Mask(err)
	}

	treeClaims, err := h.markets.Claims(nil, treeId)
	if err != nil {
		return claim, treeId, tracer.Mask(err)
	}

	var lastClaimIndex int
	{
		for i, c := range treeClaims {
			if c.Expiration.Cmp(big.NewInt(0)) == 0 {
				lastClaimIndex = i
				break
			}
		}
	}

	return treeClaims[lastClaimIndex-1], treeId, nil
}

func (h *SystemHandler) prepareVote(claim marketscontract.IMarketsClaim, treeId *big.Int) (err error) {
	// pick random voters from the available stakers
	yeaVoters, nayVoters, err := h.pickRandomVoters(claim)
	if err != nil {
		return tracer.Mask(err)
	}

	// make the actual write call to randomizer.PrepareVote()
	err = h.write(claim, treeId, yeaVoters, nayVoters)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

// @todo - make sure there's no duplicates
func (h *SystemHandler) pickRandomVoters(claim marketscontract.IMarketsClaim) (yeaVoters []string, nayVoters []string, err error) {
	yeaStakersLength := len(claim.Stake.YeaStakers)
	nayStakersLength := len(claim.Stake.NayStakers)
	if yeaStakersLength > 0 && nayStakersLength > 0 {
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
		h.log.Log(
			context.Background(),
			"level", "info",
			"message", "claimresolverhandler: invalid state: invalid stakers",
		)
		return yeaVoters, nayVoters, tracer.Mask(err)
	}

	return yeaVoters, nayVoters, nil
}

func (h *SystemHandler) write(
	claim marketscontract.IMarketsClaim,
	treeId *big.Int,
	yeaVoters []string,
	nayVoters []string,
) (err error) {
	var yeaVotersAddrs []common.Address
	{
		for _, addr := range yeaVoters {
			yeaVotersAddrs = append(yeaVotersAddrs, common.HexToAddress(addr))
		}
	}

	var nayVotersAddrs []common.Address
	{
		for _, addr := range nayVoters {
			nayVotersAddrs = append(nayVotersAddrs, common.HexToAddress(addr))
		}
	}

	privateKey, err := crypto.HexToECDSA(h.pk)
	if err != nil {
		return tracer.Mask(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return tracer.Maskf(runtime.ExecutionFailedError, "failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return tracer.Mask(err)
	}

	gasPrice, err := h.client.SuggestGasPrice(context.Background())
	if err != nil {
		return tracer.Mask(err)
	}

	chainID, ok := new(big.Int).SetString(h.cid, 10)
	if !ok {
		return tracer.Maskf(runtime.ExecutionFailedError, "failed to cast chain ID to big.Int")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return tracer.Mask(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(h.cas.Markets)
	instance, err := randomizer.NewRandomizer(address, h.client)
	if err != nil {
		return tracer.Mask(err)
	}

	tx, err := instance.PrepareVote(auth, yeaVotersAddrs, nayVotersAddrs, treeId)
	if err != nil {
		return tracer.Mask(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

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
