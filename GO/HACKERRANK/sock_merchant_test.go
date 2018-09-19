package sock_merchant

import (
	"testing"
)

func TestSockMerchant(t *testing.T) {
	tests := []struct{
		desc string
		ar []int32
		want int32
	}{
		{
			"success",
			[]int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			3,
		},
		{
			"success",
			[]int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3},
			4,
		},
	}

	for i, tc := range tests {
		if got := SockMerchant(tc.ar); got != tc.want {
			t.Errorf("TestSockMerchant(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
