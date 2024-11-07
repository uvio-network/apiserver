package notehandler

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/uvio-network/apigocode/pkg/note"
)

func Test_Server_Handler_Note_Search_Fuzz(t *testing.T) {
	var han note.API
	{
		han = tesHan()
	}

	var fuz *fuzz.Fuzzer
	{
		fuz = fuzz.New()
	}

	for i := 0; i < 1000; i++ {
		var inp *note.SearchI
		{
			inp = &note.SearchI{}
		}

		{
			fuz.Fuzz(inp)
		}

		{
			_, _ = han.Search(tesCtx(), inp)
		}
	}
}
