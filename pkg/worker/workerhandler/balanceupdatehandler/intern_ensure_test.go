package balanceupdatehandler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
)

func Test_Worker_WorkerHandler_BalanceUpdateHandler_InternEnsure_balTre(t *testing.T) {
	testCases := []struct {
		res objectid.ID
		tre []*poststorage.Object
		bal *poststorage.Object
	}{
		// Case 000
		{
			res: "1",
			tre: nil,
			bal: nil,
		},
		// Case 001
		{
			res: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
			},
			bal: nil,
		},
		// Case 002
		{
			res: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			bal: nil,
		},
		// Case 003
		{
			res: "2",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
				{ID: "3", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "2"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			bal: &poststorage.Object{
				ID: "3", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "2",
			},
		},
		// Case 004
		{
			res: "2",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "2"},
				{ID: "4", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "2"},
				{ID: "5", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			bal: &poststorage.Object{
				ID: "4", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "2",
			},
		},
		// Case 005
		{
			res: "2",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
				{ID: "3", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "2"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "2"},
				{ID: "5", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			bal: &poststorage.Object{
				ID: "3", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "2",
			},
		},
		// Case 005
		{
			res: "11",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "5", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
				{ID: "6", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "5"},

				{ID: "7", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "5"},
				{ID: "8", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "7"},
				{ID: "9", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "7"},
				{ID: "10", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "7"},
				{ID: "11", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "7"},
				{ID: "12", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "13", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "14", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "11"},

				{ID: "15", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "11"},
				{ID: "16", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "15"},
				{ID: "17", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "15"},
				{ID: "18", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "15"},
				{ID: "19", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "15"},
				{ID: "20", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "19"},
			},
			bal: &poststorage.Object{
				ID: "14", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleBalance}, Parent: "11",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			bal, err := balTre(tc.res, tc.tre)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(bal, tc.bal) {
				t.Fatalf("expected %#v got %#v", tc.bal, bal)
			}
		})
	}
}
