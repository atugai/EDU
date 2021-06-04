package hackerrank

import (
	"testing"
)

func TestPageCount(t *testing.T) {
	tests := []struct {
		desc string
		n    int32
		p    int32
		want int32
	}{
		{
			"1st page from begin",
			6,
			2,
			1,
		}, {
			"0 page from end ",
			5,
			4,
			0,
		}, {
			"0 page from start",
			5,
			5,
			0,
		},
	}

	for i, tc := range tests {
		if got := PageCount(tc.n, tc.p); got != tc.want {
			t.Errorf("PageCount(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
