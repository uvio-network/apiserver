package walletstorage

import (
	"time"

	"github.com/uvio-network/apiserver/pkg/format/walletaddress"
	"github.com/xh3b4sd/objectid"
	"github.com/xh3b4sd/tracer"
)

type Object struct {
	Address     string      `json:"address,omitempty"`
	Created     time.Time   `json:"created,omitempty"`
	Description string      `json:"description,omitempty"`
	ID          objectid.ID `json:"id,omitempty"`
	Kind        string      `json:"kind,omitempty"`
	Owner       objectid.ID `json:"owner,omitempty"`
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
