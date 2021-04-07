package hackerrank

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestCavityMap(t *testing.T) {
	tests := []struct{
		in []string
		want []string
	}{
		{
			in: []string{"12"},
			want: []string{"12"},
		},{
			in: []string{"1112", "1912", "1892", "1234"},
			want: []string{"1112", "1X12", "18X2", "1234"},
		},{
			in: []string{"999", "999", "999"},
			want: []string{"999", "999", "999"},
		},
	}

	for i, tc := range tests {
		got := CavityMap(tc.in)
		if diff := pretty.Compare(got, tc.want); diff != "" {
			t.Errorf("TestCavityMap(%d): unexpected result -got +want diff: %v", i, diff)
		}
	}
}
