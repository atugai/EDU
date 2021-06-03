package hackerrank

import (
	"testing"
)

func TestStrangeCounter(t *testing.T) {
	tests := []struct {
		in   int64
		want int64
	}{
		{
			in:   3,
			want: 1,
		}, {
			in:   8,
			want: 2,
		}, {
			in:   15,
			want: 7,
		}, {
			in:   150000,
			want: 46606,
		}, {
			in:   1230000,
			want: 342862,
		}, {
			in:   122330000,
			want: 78996590,
		},
	}

	for i, tc := range tests {
		if got := StrangeCounter(tc.in); got != tc.want {
			t.Errorf("TestStrangeCounter(%d): unexpected result got: %v want: %v", i, got, tc.want)
		}
	}
}
