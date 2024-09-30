package createresolvehandler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uvio-network/apiserver/pkg/object/objectfield"
	"github.com/uvio-network/apiserver/pkg/object/objectid"
	"github.com/uvio-network/apiserver/pkg/object/objectlabel"
	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
)

func Test_Worker_WorkerHandler_CreateResolveHandler_InternEnsure_resTre(t *testing.T) {
	testCases := []struct {
		pod objectid.ID
		tre []*poststorage.Object
		res *poststorage.Object
	}{
		// Case 000
		{
			pod: "1",
			tre: nil,
			res: nil,
		},
		// Case 001
		{
			pod: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
			},
			res: nil,
		},
		// Case 002
		{
			pod: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			res: nil,
		},
		// Case 003
		{
			pod: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			res: &poststorage.Object{
				ID: "2", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1",
			},
		},
		// Case 004
		{
			pod: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
			},
			res: &poststorage.Object{
				ID: "3", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1",
			},
		},
		// Case 005
		{
			pod: "1",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "5", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
			},
			res: &poststorage.Object{
				ID: "5", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1",
			},
		},
		// Case 005
		{
			pod: "6",
			tre: []*poststorage.Object{
				{ID: "1", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "5", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},

				{ID: "6", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "5"},
				{ID: "7", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "8", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "9", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "10", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "6"},

				{ID: "11", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "10"},
				{ID: "12", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "13", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "14", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "15", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "11"},
			},
			res: &poststorage.Object{
				ID: "10", Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "6",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			res, err := resTre(tc.pod, tc.tre)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(res, tc.res) {
				t.Fatalf("expected %#v got %#v", tc.res, res)
			}
		})
	}
}
