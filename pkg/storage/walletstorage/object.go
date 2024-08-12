package walletstorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/format/walletaddress"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Active      bool        `json:"active"`
	Address     string      `json:"address"`
	Created     time.Time   `json:"created"`
	Description string      `json:"description"`
	ID          objectid.ID `json:"id"`
	Kind        string      `json:"kind"`
	Owner       objectid.ID `json:"owner"`
	Provider    string      `json:"provider"`
}

func (o *Object) Verify() error {
	{
		if o.Address == "" {
			return tracer.Mask(WalletAddressEmptyError)
		}
		if !walletaddress.Verify(o.Address) {
			return tracer.Mask(WalletAddressFormatError)
		}
	}

	{
		if len(o.Description) > 100 {
			return tracer.Maskf(WalletDescriptionLimitError, "%d", len(o.Description))
		}
	}

	{
		if o.Kind != "contract" && o.Kind != "signer" {
			return tracer.Maskf(WalletKindInvalidError, o.Kind)
		}
	}

	{
		if o.Owner == "" {
			return tracer.Mask(WalletOwnerEmptyError)
		}
	}

	// TODO validate kind and provider combos

	{
		if o.Provider != "biconomy" && o.Provider != "external" && o.Provider != "privy" {
			return tracer.Maskf(WalletProviderInvalidError, o.Kind)
		}
	}

	return nil
}
