package hackerrank

import (
	"testing"
)

func TestDivisibleSumPairs(t *testing.T) {
	tests := []struct{
		desc string
		ar []int32
		k int32
		want int32
	}{
		{
			desc: "success",
			ar: []int32{1, 3, 2, 6, 1, 2},
			k: 3,
			want: 5,
		},
	}

	for i, tc := range tests {
		if got := DivisibleSumPairs(0, tc.k, tc.ar); got != tc.want {
			t.Errorf("TestDivisibleSumPairs(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
