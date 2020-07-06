package permutations_with_repetitions

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestPermutations(t *testing.T) {
	tests := []struct{
		desc string
		e []int
		l int
		want [][]int
	}{
		{
			desc: "empty result",
			l: 0,
			e: []int{1, 2},
			want: [][]int{},
		}, {
			desc: "different elements",
			l: 2,
			e: []int{1, 2},
			want: [][]int{
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
			},
		}, {
			desc: "similar elements",
			l: 2,
			e: []int{1, 1},
			want: [][]int{
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
			},
		}, {
			desc: "more elements",
			l: 2,
			e: []int{1, 2, 3},
			want: [][]int{
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 1},
				{2, 2},
				{2, 3},
				{3, 1},
				{3, 2},
				{3, 3},
			},
		},
	}

	for _, tc := range tests {
		got := Permutations(tc.e, tc.l)
		if diff := pretty.Compare(got, tc.want); diff != "" {
			t.Errorf("TestPermutations(%s): unexpected result -got +want diff: %v", tc.desc, diff)
		}
	}
}
