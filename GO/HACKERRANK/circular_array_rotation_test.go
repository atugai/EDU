package hackerrank

import (
	"reflect"
	"testing"
)

func TestCircularArrayRotation(t *testing.T) {
	tests := []struct{
		desc string
		slice []int32
		steps int32
		queries []int32
		want []int32
	}{
		{
			"success",
			[]int32{10, 20, 30},
			2,
			[]int32{0, 1, 2},
			[]int32{20, 30, 10},
		},
	}

	for i, tc := range tests {
		if got := CircularArrayRotation(tc.slice, tc.steps, tc.queries); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TestCircularArrayRotation(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
