package hackerrank

import "testing"

func TestFairRations(t *testing.T) {
	tests := []struct{
		in []int32
		want int32
	}{
		{
			in: []int32{2, 3, 4, 5, 6},
			want: 4,
		},{
			in: []int32{1, 2},
			want: -1,
		},
	}

	for i, tc := range tests {
		if got := FairRations(tc.in); got != tc.want {
			t.Errorf("TestFairRations(%d): unexpected result got: %v, want: %v", i, got, tc.want)
		}
	}
}
