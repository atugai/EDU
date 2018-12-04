package sequence_equation

import (
	"reflect"
	"testing"
)

func TestSequenceEquation(t *testing.T) {
	tests := []struct{
		desc string
		in []int32
		want []int32
	}{
		{
			desc: "success",
			in: []int32{2, 3, 1},
			want: []int32{2, 3, 1},
		}, {
			desc: "success",
			in: []int32{5, 2, 1, 3, 4},
			want: []int32{4, 2, 5, 1, 3},
		}, {
			desc: "success",
			in: []int32{4, 3, 5, 1, 2},
			want: []int32{1, 3, 5, 4, 2},
		},
	}

	for i, tc := range tests {
		if got := SequenceEquation(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("SequenceEquation(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
