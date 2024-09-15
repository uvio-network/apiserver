package walletstorage

import (
	"github.com/xh3b4sd/tracer"
)

func (r *Redigo) CreateWallet(inp []*Object) error {
	var err error

	for i := range inp {
		// Create the normalized key-value pair for the wallet object. With that we
		// can search for wallet objects using their IDs.
		{
			err = r.red.Simple().Create().Element(walObj(inp[i].ID), musStr(inp[i]))
			if err != nil {
				return tracer.Mask(err)
			}
		}

		// Track the user creating this wallet as the owner, and make sure that we can
		// find all wallets for any given user ID.
		{
			err = r.red.Sorted().Create().Score(walOwn(inp[i].Owner), inp[i].ID.String(), inp[i].ID.Float())
			if err != nil {
				return tracer.Mask(err)
			}
		}
	}

	return nil
}
