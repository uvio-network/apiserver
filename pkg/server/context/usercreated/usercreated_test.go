package usercreated

import (
	"context"
	"testing"
)

func Test_Context_UserCreated(t *testing.T) {
	{
		v := FromContext(context.Background())
		if v != false {
			t.Fatal("user must be empty")
		}
	}

	{
		v := FromContext(NewContext(context.Background(), true))
		if v != true {
			t.Fatal("user must be test")
		}
	}
}
