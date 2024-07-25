package posthandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Delete(ctx context.Context, req *post.DeleteI) (*post.DeleteO, error) {
	fmt.Printf("/post.API/Delete not implemented\n")
	return &post.DeleteO{}, nil
}
