package wallethandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/wallet"
)

func (h *Handler) Update(ctx context.Context, req *wallet.UpdateI) (*wallet.UpdateO, error) {
	fmt.Printf("/wallet.API/Update not implemented\n")
	return &wallet.UpdateO{}, nil
}
