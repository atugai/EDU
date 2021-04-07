package hackerrank

import (
	"testing"
)

func TestGetTotal(t *testing.T) {
	tests := []struct{
		desc string
		a []int32
		b []int32
		want int32
	}{
		{
			desc: "success",
			a: []int32{2,4},
			b: []int32{16,32,96},
			want: int32(3),
		},
		{
			desc: "success",
			a: []int32{3,4},
			b: []int32{24,48},
			want: int32(2),
		},
	}

	for i, tc := range tests {
		if got := GetTotal(tc.a, tc.b); got != tc.want {
			t.Errorf("TestGetTotal(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
