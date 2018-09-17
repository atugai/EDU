package migratory_birds

import (
	"testing"
)

func TestMigratoryBirds(t *testing.T) {
	tests := []struct{
		desc string
		arr []int32
		want int32
	}{
		{
			"success",
			[]int32{1, 4, 4, 4, 5, 3},
			4,
		},
		{
			"success",
			[]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4},
			3,
		},
	}

	for i, tc := range tests {
		if got := MigratoryBirds(tc.arr); got != tc.want {
			t.Errorf("MigratoryBirds(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}

}
