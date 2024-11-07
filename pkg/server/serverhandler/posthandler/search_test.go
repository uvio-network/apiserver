package posthandler

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/post"
)

func Test_Server_Handler_Post_Search_Fuzz(t *testing.T) {
	var han post.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *post.SearchI
		{
			inp = &post.SearchI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Search(tesCtx(), inp)
		}
	}
}
