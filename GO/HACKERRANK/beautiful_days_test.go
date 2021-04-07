package hackerrank

import (
	"testing"
)

func TestBeautifulDays(t *testing.T) {
	tests := []struct{
		desc string
		i int32
		j int32
		k int32
		want int32
	}{
		{
			"success",
			20,
			23,
			6,
			2,
		},
		{
			"success",
			13,
			45,
			3,
			33,
		},
	}

	for i, tc := range tests {
		if got := BeautifulDays(tc.i, tc.j, tc.k); got != tc.want {
			t.Errorf("TestBeautifulDays(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
