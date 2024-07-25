package userhandler

import (
	"github.com/gorilla/mux"
	"github.com/uvio-network/apigocode/pkg/user"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := user.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
