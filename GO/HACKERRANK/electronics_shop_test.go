package electronics_shop

import (
	"testing"
)

func TestGetMoneySpent(t *testing.T) {
	tests := []struct{
		desc string
		kb []int32
		dr []int32
		b int32
		want int32
	}{
		{
			"success shopping",
			[]int32{3, 1},
			[]int32{5,2,8},
			10,
			9,
		},
		{
			"no shopping possible",
			[]int32{4},
			[]int32{5},
			5,
			-1,
		},
	}

	for i, tc := range tests {
		if got := GetMoneySpent(tc.kb, tc.dr, tc.b); got != tc.want {
			t.Errorf("TestGetMoneySpent(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
