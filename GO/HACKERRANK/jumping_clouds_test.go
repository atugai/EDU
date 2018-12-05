package jumping_clouds

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct{
		desc string
		clouds []int32
		step int32
		want int32
	}{
		{
			desc: "success",
			clouds: []int32{0, 0, 1, 0},
			step: 2,
			want: 96,
		}, {
			desc: "success",
			clouds: []int32{0, 0, 1, 0, 0, 1, 1, 0},
			step: 2,
			want: 92,
		}, {
			desc: "success",
			clouds: []int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0},
			step: 3,
			want: 94,
		},
	}

	for i, tc := range tests {
        	if got := JumpingOnClouds(tc.clouds, tc.step); got != tc.want {
        	        t.Errorf("JumpingOnClouds(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
        	}
        }
}
