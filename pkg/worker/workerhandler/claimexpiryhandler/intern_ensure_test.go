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

func Test_Worker_WorkerHandler_ClaimExpiryHandler_InternEnsure_finTre(t *testing.T) {
	testCases := []struct {
		tim time.Time
		tre []*poststorage.Object
		res *poststorage.Object
		fin bool
	}{
		// Case 000
		{
			tim: time.Unix(1, 0),
			tre: nil,
			res: nil,
			fin: false,
		},
		// Case 001
		{
			tim: time.Unix(1, 0),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
			},
			res: nil,
			fin: false,
		},
		// Case 002
		{
			tim: time.Unix(1, 0),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			res: nil,
			fin: false,
		},
		// Case 003
		{
			tim: time.Unix(1, 0),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Expiry: time.Unix(5, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
			},
			res: nil,
			fin: false,
		},
		// Case 004
		{
			tim: time.Unix(5, 0).Add(oneWeek),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Expiry: time.Unix(5, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
			},
			res: nil,
			fin: false,
		},
		// Case 005, same as above but after the challenge period
		{
			tim: time.Unix(6, 0).Add(oneWeek),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Expiry: time.Unix(5, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
			},
			res: &poststorage.Object{
				ID: "3", Expiry: time.Unix(5, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1",
			},
			fin: true,
		},
		// Case 006
		{
			tim: time.Unix(1, 0),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(1, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "5", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},
			},
			res: nil,
			fin: false,
		},
		// Case 007
		{
			tim: time.Unix(11, 0),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(1, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "5", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},

				{ID: "6", Expiry: time.Unix(5, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "5"},
				{ID: "7", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "8", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "9", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "10", Expiry: time.Unix(7, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "6"},

				{ID: "11", Expiry: time.Unix(9, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "10"},
				{ID: "12", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "13", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "14", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "15", Expiry: time.Unix(11, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "11"},
			},
			res: nil,
			fin: false,
		},
		// Case 008, same as above but after the expiry of the final resolve, no
		// challenge window needed
		{
			tim: time.Unix(12, 0),
			tre: []*poststorage.Object{
				{ID: "1", Expiry: time.Unix(1, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecyclePropose}, Parent: ""},
				{ID: "2", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "3", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "4", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "1"},
				{ID: "5", Expiry: time.Unix(3, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "1"},

				{ID: "6", Expiry: time.Unix(5, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "5"},
				{ID: "7", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "8", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "9", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "6"},
				{ID: "10", Expiry: time.Unix(7, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "6"},

				{ID: "11", Expiry: time.Unix(9, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleDispute}, Parent: "10"},
				{ID: "12", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "13", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "14", Kind: "comment", Lifecycle: objectfield.Lifecycle{}, Parent: "11"},
				{ID: "15", Expiry: time.Unix(11, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "11"},
			},
			res: &poststorage.Object{
				ID: "15", Expiry: time.Unix(11, 0), Kind: "claim", Lifecycle: objectfield.Lifecycle{Data: objectlabel.LifecycleResolve}, Parent: "11",
			},
			fin: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			res, fin := finTre(tc.tim, tc.tre)

			if !reflect.DeepEqual(res, tc.res) {
				t.Fatalf("expected res %#v got %#v", tc.res, res)
			}
			if fin != tc.fin {
				t.Fatalf("expected fin %#v got %#v", tc.fin, fin)
			}
		})
	}
}
