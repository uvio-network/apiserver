package claimresolvehandler

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uvio-network/apiserver/pkg/contract/marketscontract"
	"github.com/uvio-network/apiserver/pkg/contract/randomizercontract"
	"github.com/uvio-network/apiserver/pkg/storage"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type SystemHandlerConfig struct {
	Cid string
	Pk  string
	Log logger.Interface
	Rpc string
	Sto storage.Interface
	Cas ContractAddresses
}

type ContractAddresses struct {
	Markets    string
	Randomizer string
}

type SystemHandler struct {
	log        logger.Interface
	pk         string
	markets    *marketscontract.Markets
	randomizer *randomizercontract.Randomizer
	client     *ethclient.Client
	sto        storage.Interface
	Cas        ContractAddresses
}

func NewSystemHandler(c SystemHandlerConfig) *SystemHandler {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Sto == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Sto must not be empty", c)))
	}

	var err error

	var eth *ethclient.Client
	{
		eth, err = ethclient.Dial(c.Rpc)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var mar *marketscontract.Markets
	{
		mar, err = marketscontract.NewMarkets(common.HexToAddress(c.Cas.Markets), eth)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var ran *randomizercontract.Randomizer
	{
		ran, err = randomizercontract.NewRandomizer(common.HexToAddress(c.Cas.Randomizer), eth)
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	return &SystemHandler{
		log:        c.Log,
		pk:         c.Pk,
		markets:    mar,
		randomizer: ran,
		client:     eth,
		sto:        c.Sto,
		Cas:        c.Cas,
	}
}
