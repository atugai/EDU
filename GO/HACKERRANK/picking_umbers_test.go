package picking_umbers

import (
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	tests := []struct{
		desc string
		in []int32
		want int32
	}{
		{
			"success",
			[]int32{4, 6, 5, 3, 3, 1},
			3,
		},
		{
			"success",
			[]int32{1, 2, 2, 3, 1, 2},
			5,
		},
	}

	for i, tc := range tests {
		if got := PickingNumbers(tc.in); got != tc.want {
			t.Errorf("PickingNumbers(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
