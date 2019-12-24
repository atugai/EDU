package minimum_distances

import "testing"

func TestMinimumDistance(t *testing.T) {
	tests := []struct{
		in []int32
		want int32
	}{
		{
			in: []int32{7, 1, 3, 4, 1, 7},
			want: 3,
		},{
			in: []int32{1, 2},
			want: -1,
		},
	}

	for i, tc := range tests {
		if got := MinimumDistance(tc.in); got != tc.want {
			t.Errorf("TestMinimumDistance(%d): unexpected result got: %v, want: %v", i, got, tc.want)
		}
	}
}
