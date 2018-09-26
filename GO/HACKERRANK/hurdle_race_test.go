package hurdle_race

import (
	"testing"
)

func TestHurdleRace(t *testing.T) {
	tests := []struct{
		desc string
		height []int32
		k int32
		want int32
	}{
		{
			"need potions",
			[]int32{1, 6, 3, 5, 2},
			4,
			2,
		},
		{
			"no potions needed",
			[]int32{2, 5, 4, 5, 2},
			7,
			0,
		},
	}

	for i, tc := range tests {
		if got := HurdleRace(tc.k, tc.height); got != tc.want {
			t.Errorf("HurdleRace(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
