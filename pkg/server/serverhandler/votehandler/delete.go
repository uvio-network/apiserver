package votehandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/vote"
)

func (h *Handler) Delete(ctx context.Context, req *vote.DeleteI) (*vote.DeleteO, error) {
	fmt.Printf("/vote.API/Delete not implemented\n")
	return &vote.DeleteO{}, nil
}
