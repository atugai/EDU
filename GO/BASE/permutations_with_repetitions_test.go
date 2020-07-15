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
		f bool
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
		}, {
			desc: "full permutations",
			l: 2,
			e: []int{1, 2, 3},
			f: true,
			want: [][]int{
				{1},
				{2},
				{3},
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
			desc: "looong permutations pointer test",
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
		got := Permutations(tc.e, tc.l, tc.f)
		if diff := pretty.Compare(got, tc.want); diff != "" {
			t.Errorf("TestPermutations(%s): unexpected result -got +want diff: %v", tc.desc, diff)
		}
	}
}

func BenchmarkPermutations(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Permutations([]int{1,2,3}, 10, false)
  }
}
