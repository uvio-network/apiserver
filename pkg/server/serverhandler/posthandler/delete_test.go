package posthandler

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/post"
)

func Test_Server_Handler_post_Delete_Fuzz(t *testing.T) {
	var han post.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *post.DeleteI
		{
			inp = &post.DeleteI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Delete(tesCtx(), inp)
		}
	}
}
