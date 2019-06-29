package taum_and_bday

import "testing"

func TestMinimumPrice(t *testing.T) {
	tests := []struct{
		b int64
		w int64
		bc int64
		wc int64
		z int64
		want int64
	}{
		{
			b: 10,
			w: 10,
			bc: 1,
			wc: 1,
			z: 1,
			want: 20,
		},{
			b: 5,
			w: 9,
			bc: 2,
			wc: 3,
			z: 4,
			want: 37,
		},{
			b: 3,
			w: 6,
			bc: 9,
			wc: 1,
			z: 1,
			want: 12,
		},{
			b: 7,
			w: 7,
			bc: 4,
			wc: 2,
			z: 1,
			want: 35,
		},{
			b: 3,
			w: 3,
			bc: 1,
			wc: 9,
			z: 2,
			want: 12,
		},{
			b: 3,
			w: 5,
			bc: 3,
			wc: 4,
			z: 1,
			want: 29,
		},
	}

	for i, tc := range tests {
		if got := MinimumPrice(tc.b, tc.w, tc.bc, tc.wc, tc.z); got != tc.want {
			t.Errorf("TestMinimumPrice(%d): unexpected result got: %v, want: %v", i, got, tc.want)
		}
	}
}
