package kaprekar_numbers

import( 
	"testing"
	"reflect"
)

func TestMinimumPrice(t *testing.T) {
        tests := []struct{
                p int32
                q int32
                want []int32
        }{
		{
			p: 1,
			q: 99999,
			want: []int32{1, 9, 45, 55, 99, 297, 703, 999, 2223, 2728, 4950, 5050, 7272, 7777, 9999, 17344, 22222, 77778, 82656, 95121, 99999},
		}, {
			p: 400,
			q: 700,
			want: nil,
		},
	}

        for i, tc := range tests {
                if got := KaprekarNumbers(tc.p, tc.q); !reflect.DeepEqual(got, tc.want) {
                        t.Errorf("TestMinimumPrice(%d): unexpected result got: %v, want: %v", i, got, tc.want)
                }
        }
}
