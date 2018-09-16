package breaking_records

import (
	"reflect"
	"testing"
)

func TestBreakingRecords(t *testing.T) {
	tests := []struct{
		desc string
		scores []int32
		want []int32

	}{
		{
			desc: "success",
			scores: []int32{10,5,20,20,4,5,2,25,1},
			want: []int32{2,4},
		},
		{
			desc: "success",
			scores: []int32{3,4,21,36,10,28,35,5,24,42},
			want: []int32{4,0},
		},
	}

	for i, tc := range tests {
		if got := BreakingRecords(tc.scores); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TestBreakingRecords(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
