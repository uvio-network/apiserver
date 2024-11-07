package notehandler

import (
	"github.com/gorilla/mux"
	"github.com/uvio-network/apigocode/pkg/note"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := note.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
