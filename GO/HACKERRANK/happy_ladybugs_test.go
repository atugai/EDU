package hackerrank

import "testing"

func TestHappyLadybugs(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "_",
			want: true,
		}, {
			in:   "__",
			want: true,
		}, {
			in:   "AABBC_C",
			want: true,
		}, {
			in:   "DD__FQ_QQF",
			want: true,
		}, {
			in:   "RBY_YBR",
			want: true,
		}, {
			in:   "B_RRBR",
			want: true,
		}, {
			in:   "AABBCC",
			want: true,
		}, {
			in:   "AAABB",
			want: true,
		}, {
			in:   "A",
			want: false,
		}, {
			in:   "AABBC",
			want: false,
		}, {
			in:   "AABCBC",
			want: false,
		}, {
			in:   "X_Y__X",
			want: false,
		},
	}

	for _, tc := range tests {
		if got := HappyLadybugs(tc.in); got != tc.want {
			t.Errorf("HappyLadybugs(%q): unexpected result got: %v, want: %v", tc.in, got, tc.want)
		}
	}
}
