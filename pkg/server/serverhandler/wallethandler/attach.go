package wallethandler

import (
	"github.com/gorilla/mux"
	"github.com/uvio-network/apigocode/pkg/wallet"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := wallet.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
