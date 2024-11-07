package notehandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/note"
)

func (h *Handler) Create(ctx context.Context, req *note.CreateI) (*note.CreateO, error) {
	fmt.Printf("/note.API/Create not implemented\n")
	return &note.CreateO{}, nil
}
