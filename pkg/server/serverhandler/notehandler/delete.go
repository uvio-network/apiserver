package notehandler

import (
	"context"
	"fmt"

	"github.com/uvio-network/apigocode/pkg/note"
)

func (h *Handler) Delete(ctx context.Context, req *note.DeleteI) (*note.DeleteO, error) {
	fmt.Printf("/note.API/Delete not implemented\n")
	return &note.DeleteO{}, nil
}
