package postreconciler

import (
	"fmt"
	"slices"
	"testing"

	"github.com/uvio-network/apiserver/pkg/storage/poststorage"
	"github.com/uvio-network/apiserver/pkg/storage/votestorage"
)

func Test_Reconciler_PostReconciler_updateVotes_claim(t *testing.T) {
	testCases := []struct {
		pos []*poststorage.Object
		vot []*votestorage.Object
		out [][]float64
	}{
		// Case 000
		{
			pos: []*poststorage.Object{
				{Kind: "claim", Owner: "1", Votes: []float64{0, 0, 0, 0}},
			},
			vot: []*votestorage.Object{
				{Option: true, Owner: "1", Value: 1},
			},
			out: [][]float64{
				{1, 0, 1, 1},
			},
		},
		// Case 001
		{
			pos: []*poststorage.Object{
				{Kind: "claim", Owner: "1", Votes: []float64{1, 0, 1, 1}},
			},
			vot: []*votestorage.Object{
				{Option: true, Owner: "1", Value: 1},
			},
			out: [][]float64{
				{2, 0, 1, 2},
			},
		},
		// Case 002
		{
			pos: []*poststorage.Object{
				{Kind: "claim", Owner: "1", Votes: []float64{2, 0, 1, 2}},
			},
			vot: []*votestorage.Object{
				{Option: false, Owner: "2", Value: 3},
			},
			out: [][]float64{
				{2, 3, 1, 2},
			},
		},
		// Case 003
		{
			pos: []*poststorage.Object{
				{Kind: "claim", Owner: "1", Votes: []float64{2, 3, 1, 2}},
			},
			vot: []*votestorage.Object{
				{Option: true, Owner: "3", Value: 1.5},
			},
			out: [][]float64{
				{3.5, 3, 1, 2},
			},
		},
		// Case 004
		{
			pos: []*poststorage.Object{
				{Kind: "claim", Owner: "1", Votes: []float64{0, 0, 0, 0}},
				{Kind: "claim", Owner: "1", Votes: []float64{1, 0, 1, 1}},
				{Kind: "claim", Owner: "1", Votes: []float64{2, 0, 1, 2}},
				{Kind: "claim", Owner: "1", Votes: []float64{2, 3, 1, 2}},
			},
			vot: []*votestorage.Object{
				{Option: true, Owner: "1", Value: 1},
				{Option: true, Owner: "1", Value: 1},
				{Option: false, Owner: "2", Value: 3},
				{Option: true, Owner: "3", Value: 1.5},
			},
			out: [][]float64{
				{1, 0, 1, 1},
				{2, 0, 1, 2},
				{2, 3, 1, 2},
				{3.5, 3, 1, 2},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			if len(tc.pos) != len(tc.vot) {
				t.Fatalf("for every post object there must be one vote object")
			}

			var out [][]float64
			for i := range tc.pos {
				out = append(out, updateVotes(tc.pos[i], tc.vot[i]).Votes)
			}

			if len(out) != len(tc.out) {
				t.Fatalf("expected %#v got %#v", len(tc.out), len(out))
			}

			for i := range out {
				if !slices.Equal(out[i], tc.out[i]) {
					t.Fatalf("[%d] expected %#v got %#v", i, tc.out[i], out[i])
				}
			}
		})
	}
}

func Test_Reconciler_PostReconciler_updateVotes_comment(t *testing.T) {
	testCases := []struct {
		pos []*poststorage.Object
		vot []*votestorage.Object
		out [][]float64
	}{
		// Case 000
		{
			pos: []*poststorage.Object{
				{Kind: "comment", Votes: []float64{0, 0}},
			},
			vot: []*votestorage.Object{
				{Option: true, Value: 1},
			},
			out: [][]float64{
				{1, 0},
			},
		},
		// Case 001
		{
			pos: []*poststorage.Object{
				{Kind: "comment", Votes: []float64{1, 0}},
			},
			vot: []*votestorage.Object{
				{Option: true, Value: 1},
			},
			out: [][]float64{
				{2, 0},
			},
		},
		// Case 002
		{
			pos: []*poststorage.Object{
				{Kind: "comment", Votes: []float64{2, 0}},
			},
			vot: []*votestorage.Object{
				{Option: false, Value: 3},
			},
			out: [][]float64{
				{2, 3},
			},
		},
		// Case 003
		{
			pos: []*poststorage.Object{
				{Kind: "comment", Votes: []float64{2, 3}},
			},
			vot: []*votestorage.Object{
				{Option: true, Value: 1.5},
			},
			out: [][]float64{
				{3.5, 3},
			},
		},
		// Case 003
		{
			pos: []*poststorage.Object{
				{Kind: "comment", Votes: []float64{0, 0}},
				{Kind: "comment", Votes: []float64{1, 0}},
				{Kind: "comment", Votes: []float64{2, 0}},
				{Kind: "comment", Votes: []float64{2, 3}},
			},
			vot: []*votestorage.Object{
				{Option: true, Value: 1},
				{Option: true, Value: 1},
				{Option: false, Value: 3},
				{Option: true, Value: 1.5},
			},
			out: [][]float64{
				{1, 0},
				{2, 0},
				{2, 3},
				{3.5, 3},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			if len(tc.pos) != len(tc.vot) {
				t.Fatalf("for every post object there must be one vote object")
			}

			var out [][]float64
			for i := range tc.pos {
				out = append(out, updateVotes(tc.pos[i], tc.vot[i]).Votes)
			}

			if len(out) != len(tc.out) {
				t.Fatalf("expected %#v got %#v", len(tc.out), len(out))
			}

			for i := range out {
				if !slices.Equal(out[i], tc.out[i]) {
					t.Fatalf("[%d] expected %#v got %#v", i, tc.out[i], out[i])
				}
			}
		})
	}
}
