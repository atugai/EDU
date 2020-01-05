package cavity_map

import (
	"testing"
	"reflect"
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
		if got := CavityMap(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TestCavityMap(%d): unexpected result got: %v, want: %v", i, got, tc.want)
		}
	}
}
