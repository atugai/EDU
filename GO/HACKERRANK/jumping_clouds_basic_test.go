package hackerrank

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
        tests := []struct{
                desc string
                c []int32
                want int32
        }{
		{
			desc: "success jump",
			c: []int32{0, 0, 1, 0, 0, 1, 0},
			want: 4,
		},
		{
			desc: "success jump",
			c: []int32{0, 0, 0, 1, 0, 0},
			want: 3,
		},
	}

	for i, tc := range tests {
		if got := JumpingOnClouds(tc.c); got != tc.want {
			t.Errorf("JumpingOnClouds(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
