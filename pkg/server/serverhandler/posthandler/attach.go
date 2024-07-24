package posthandler

import (
	"github.com/gorilla/mux"
	"github.com/uvio-network/apigocode/pkg/post"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := post.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
