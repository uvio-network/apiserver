package userhandler

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/user"
)

func Test_Server_Handler_User_Update_Fuzz(t *testing.T) {
	var han user.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *user.UpdateI
		{
			inp = &user.UpdateI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Update(tesCtx(), inp)
		}
	}
}
