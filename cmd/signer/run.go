package signer

import (
	"fmt"
	// "os"

	"github.com/spf13/cobra"
	// "github.com/uvio-network/apiserver/pkg/runtime"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	fmt.Println("Hello, World!")
	// r.generateWallet()
	// r.generateKeystore()
	r.sendTransaction()
	// r.loadKeystore()
	fmt.Println("Hello, World2!")
}
