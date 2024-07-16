package hackerrank

import (
	"reflect"
	"testing"
)

func TestBigSorting(t *testing.T) {
	tests := []struct {
		desc string
		in   []string
		want []string
	}{
		{
			desc: "TestCase 0",
			in: []string{
				"6",
				"31415926535897932384626433832795",
				"1",
				"3",
				"10",
				"3",
			},
			want: []string{
				"1",
				"3",
				"3",
				"6",
				"10",
				"31415926535897932384626433832795",
			},
		},
		{
			desc: "TestCase 1",
			in: []string{
				"100",
				"1",
				"12303479849857341718340192371",
				"3084193741082938",
				"3084193741082937",
				"111",
				"200",
			},
			want: []string{
				"1",
				"100",
				"111",
				"200",
				"3084193741082937",
				"3084193741082938",
				"12303479849857341718340192371",
			},
		},
	}

	for i, tc := range tests {
		if got := BigSorting(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("BigSorting(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
