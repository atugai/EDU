package repeated_string

import (
	"testing"
)

func TestRepeatedString(t *testing.T) {
	tests := []struct{
		desc string
		n int64
		s string
		want int64
	}{
		{
			desc: "simple case",
			n: 10,
			s: "abcac",
			want: 4,
		}, {
			desc: "another simple case",
			n: 10,
			s: "aba",
			want: 7,
		}, {
			desc: "repeated 1 symbol",
			n: 1000000000000,
			s: "a",
			want: 1000000000000,
		}, {
			desc: "complex case",
			n: 736778906400,
			s: "kmretasscityylpdhuwjirnqimlkcgxubxmsxpypgzxtenweirknjtasxtvxemtwxuarabssvqdnktqadhyktagjxoanknhgilnm",
			want: 51574523448,
		},
	}

	for i, tc := range tests {
                if got := RepeatedString(tc.s, tc.n); got != tc.want {
                        t.Errorf("RepeatedString(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
                }
        }
}
