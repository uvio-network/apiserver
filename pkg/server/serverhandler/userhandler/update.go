package userhandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/user"
)

func (h *Handler) Update(ctx context.Context, req *user.UpdateI) (*user.UpdateO, error) {
	fmt.Printf("/user.API/Update not implemented\n")
	return &user.UpdateO{}, nil
}
