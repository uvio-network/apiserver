package posthandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Update(ctx context.Context, req *post.UpdateI) (*post.UpdateO, error) {
	fmt.Printf("/post.API/Update not implemented\n")
	return &post.UpdateO{}, nil
}
