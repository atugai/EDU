package viral_advertising

import (
	"testing"
)

func TestViralAdvertising(t *testing.T) {
	tests := []struct{
		desc string
		day int32
		want int32
	}{
		{
			"succes",
			3,
			9,
		},
		{
			"succes",
			4,
			15,
		},
	}

	for i, tc := range tests {
		if got := ViralAdvertising(tc.day); got != tc.want {
			t.Errorf("TestViralAdvertising(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
