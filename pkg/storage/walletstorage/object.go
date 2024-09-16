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
		if o.Kind != "embedded" && o.Kind != "injected" {
			return tracer.Maskf(WalletKindInvalidError, o.Kind)
		}
	}

	{
		if o.Owner == "" {
			return tracer.Mask(WalletOwnerEmptyError)
		}
	}

	return nil
}
