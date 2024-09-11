package subjectclaim

import (
	"context"
	"testing"
)

func Test_Context_SubjectClaim(t *testing.T) {
	{
		v := FromContext(context.Background())
		if v != "" {
			t.Fatal("user must be empty")
		}
	}

	{
		v := FromContext(NewContext(context.Background(), "test"))
		if v != "test" {
			t.Fatal("user must be test")
		}
	}
}
