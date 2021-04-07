package hackerrank

import (
	"testing"
	"reflect"
)

func TestAcmTeam(t *testing.T) {
	tests := []struct{
		desc string
		in []string
		want []int32
	}{
		{
			desc: "success",
			in: []string{"10101", "11100", "11010", "00101"},
			want: []int32{5, 2},
		}, {
			desc: "success",
			in: []string{"11101", "10101","11001", "10111", "10000", "01110"},
			want: []int32{5, 6},

		},
	}

        for i, tc := range tests {
                if got := AcmTeam(tc.in); !reflect.DeepEqual(got, tc.want) {
                        t.Errorf("AcmTeam(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
                }
        }
}
