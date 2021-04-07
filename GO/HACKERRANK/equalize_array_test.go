package hackerrank

import (
	"testing"
)

func TestEqualizeArray(t *testing.T) {
	tests := []struct{
		desc string
		arr []int32
		want int32
	}{
		{
			desc: "success",
			arr: []int32{3, 3, 2, 1, 3},
			want: 2,
		}, {
			desc: "success",
			arr: []int32{1, 2, 3, 1, 2, 3, 3, 3},
			want: 4,
		},
	}

        for i, tc := range tests {
                if got := EqualizeArray(tc.arr); got != tc.want {
                        t.Errorf("EqualizeArray(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
                }
        }
}
