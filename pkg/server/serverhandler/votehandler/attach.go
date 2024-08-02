package votehandler

import (
	"github.com/gorilla/mux"
	"github.com/uvio-network/apigocode/pkg/vote"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := vote.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
