package hackerrank

import "testing"

func TestHowManyGames(t *testing.T) {
	tests := []struct{
		p int32
		d int32
		m int32
		s int32
		want int32
	}{
		{
			p: 20,
			d: 3,
			m: 6,
			s: 80,
			want: 6,
		},{
			p: 20,
			d: 3,
			m: 6,
			s: 85,
			want: 7,
		},
	}

	for i, tc := range tests {
		if got := HowManyGames(tc.p, tc.d, tc.m, tc.s); got != tc.want {
			t.Errorf("TestHowManyGames(%d): unexpected result got: %v, want: %v", i, got, tc.want)
		}
	}
}
