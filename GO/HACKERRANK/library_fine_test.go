package hackerrank

import (
	"testing"
)

func TestLibraryFine(t *testing.T) {
	tests := []struct {
		desc string
		rY   int32
		rM   int32
		rD   int32
		eY   int32
		eM   int32
		eD   int32
		want int32
	}{
		{
			desc: "no fine",
			rY:   2015,
			rM:   6,
			rD:   6,
			eY:   2016,
			eM:   6,
			eD:   9,
			want: 0,
		}, {
			desc: "3 days delay",
			rY:   2015,
			rM:   6,
			rD:   9,
			eY:   2015,
			eM:   6,
			eD:   6,
			want: 45,
		}, {
			desc: "9 days delay",
			rY:   2018,
			rM:   7,
			rD:   14,
			eY:   2018,
			eM:   7,
			eD:   5,
			want: 135,
		}, {
			desc: "2 days delay",
			rY:   2018,
			rM:   9,
			rD:   14,
			eY:   2018,
			eM:   7,
			eD:   5,
			want: 1000,
		}, {
			desc: "2 days delay",
			rY:   2018,
			rM:   5,
			rD:   14,
			eY:   2018,
			eM:   7,
			eD:   5,
			want: 0,
		}, {
			desc: "2 days delay",
			rY:   2019,
			rM:   9,
			rD:   14,
			eY:   2018,
			eM:   7,
			eD:   5,
			want: 10000,
		}, {
			desc: "2 days delay",
			rY:   2019,
			rM:   9,
			rD:   14,
			eY:   2019,
			eM:   9,
			eD:   14,
			want: 0,
		},
	}

	for i, tc := range tests {
		if got := LibraryFine(tc.rD, tc.rM, tc.rY, tc.eD, tc.eM, tc.eY); got != tc.want {
			t.Errorf("LibraryFine(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
