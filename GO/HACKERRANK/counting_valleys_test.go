package counting_valleys

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	tests := []struct{
		desc string
		path string
		want int32
	}{
		{
			"one valley",
			"UDDDUDUU",
			1,
		},
		{
			"two valleys",
			"DDUUDDUDUUUD",
			2,
		},
	}

	for i, tc := range tests {
		if got := CountingValleys(tc.path); got != tc.want {
			t.Errorf("TestCountingValleys(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
