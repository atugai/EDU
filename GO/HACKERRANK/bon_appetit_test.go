package bon_appetit

import (
	"testing"
)

func TestBonAppetit(t *testing.T) {
	tests := []struct{
		desc string
		bill []int32
		excl int32
		paid int32
		want string
	}{
		{
			"bon appetit",
			[]int32{3,10,2,9},
			1,
			7,
			"Bon Appetit",
		},{
			"cash back",
			[]int32{3,10,2,9},
			1,
			12,
			"5",
		},

	}

	for i, tc := range tests {
		if got := BonAppetit(tc.bill, tc.excl, tc.paid); got != tc.want {
			t.Errorf("TestBonAppetit(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
