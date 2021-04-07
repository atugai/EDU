package hackerrank

import (
	"testing"
)

func TestLibraryFine(t *testing.T) {
	tests := []struct{
		desc string
		rY int32
		rM int32
		rD int32
		eY int32
		eM int32
		eD int32
		want int32
	}{
		{
			desc: "no fine",
			rY: 2015,
			rM: 6,
			rD: 6,
			eY: 2016,
			eM: 6,
			eD: 9,
			want: 0,
		},
		{
			desc: "3 days delay",
			rY: 2015,
			rM: 6,
			rD: 9,
			eY: 2015,
			eM: 6,
			eD: 6,
			want: 45,
		},
	}

	for i, tc := range tests {
		if got := LibraryFine(tc.rD, tc.rM, tc.rY, tc.eD, tc.eM, tc.eY); got != tc.want {
			t.Errorf("LibraryFine(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
