package permutations_with_repetitions

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestPermutationsImproved(t *testing.T) {
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
			want: [][]int{{}},
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
		}, {
			desc: "looong permutations",
			l: 4,
			e: []int{9, 11},
			want: [][]int{
				{9, 9, 9, 9},
				{9, 9, 9, 11},
				{9, 9, 11, 9},
				{9, 9, 11, 11},
				{9, 11, 9, 9},
				{9, 11, 9, 11},
				{9, 11, 11, 9},
				{9, 11, 11, 11},
				{11, 9, 9, 9},
				{11, 9, 9, 11},
				{11, 9, 11, 9},
				{11, 9, 11, 11},
				{11, 11, 9, 9},
				{11, 11, 9, 11},
				{11, 11, 11, 9},
				{11, 11, 11, 11},
			},
		},
	}

	for _, tc := range tests {
		g := NewGenerator(tc.l, tc.e)

		for i, v := range tc.want {
			got, err := g.Next()
			if err != nil {
				t.Errorf("got unexpected error(%v), want error(nil)", err)
			}

			if diff := pretty.Compare(got, v); diff != "" {
				t.Errorf("TestPermutationsImproved(%s): unexpected result for outcome(%d) -got +want diff: %v", tc.desc, i, diff)
			}
		}

		if _, err := g.Next(); err == nil {
			t.Errorf("TestPermutationsImproved(%s): unexpected nil error, want out of range", tc.desc)
		}
	}
}
