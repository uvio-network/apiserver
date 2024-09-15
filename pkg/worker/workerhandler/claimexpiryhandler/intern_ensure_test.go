package claimexpiryhandler

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
)

func Test_Worker_WorkerHandler_ClaimExpiryHandler_InternEnsure_searchResolve(t *testing.T) {
	testCases := []struct {
		tre []*poststorage.Object
		res []*poststorage.Object
	}{
		// Case 000
		{
			tre: nil,
			res: nil,
		},
		// Case 001
		{
			tre: []*poststorage.Object{
				{Expiry: time.Unix(101, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}},
			},
			res: nil,
		},
		// Case 002
		{
			tre: []*poststorage.Object{
				{Expiry: time.Unix(101, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}},
				{Expiry: time.Unix(102, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
			},
			res: []*poststorage.Object{
				{Expiry: time.Unix(102, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
			},
		},
		// Case 003
		{
			tre: []*poststorage.Object{
				{Expiry: time.Unix(101, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}},
				{Expiry: time.Unix(102, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
				{Expiry: time.Unix(103, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}},
			},
			res: nil,
		},
		// Case 004
		{
			tre: []*poststorage.Object{
				{Expiry: time.Unix(101, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}},
				{Expiry: time.Unix(102, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
				{Expiry: time.Unix(103, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}},
				{Expiry: time.Unix(104, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
			},
			res: []*poststorage.Object{
				{Expiry: time.Unix(102, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
				{Expiry: time.Unix(104, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
			},
		},
		// Case 005
		{
			tre: []*poststorage.Object{
				{Expiry: time.Unix(101, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}},
				{Expiry: time.Unix(102, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
				{Expiry: time.Unix(103, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}},
				{Expiry: time.Unix(104, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}},
				{Expiry: time.Unix(105, 0), Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}},
			},
			res: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			res := searchResolve(tc.tre)

			if !reflect.DeepEqual(res, tc.res) {
				t.Fatalf("expected %#v got %#v", tc.res, res)
			}
		})
	}
}
