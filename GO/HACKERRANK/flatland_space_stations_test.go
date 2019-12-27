package flatland_space_stations

import "testing"

func TestFlatlandSpaceStations(t *testing.T) {
	tests := []struct{
		c int32
		n []int32
		want int32
	}{
		{
			c: 5,
			n: []int32{0, 4},
			want: 2,
		},{
			c: 6,
			n: []int32{0, 1, 2, 3, 4, 5, 6},
			want: 0,
		},
	}

	for i, tc := range tests {
		if got := FlatlandSpaceStations(tc.c, tc.n); got != tc.want {
			t.Errorf("TestFlatlandSpaceStations(%d): unexpected result got: %v, want: %v", i, got, tc.want)
		}
	}
}
