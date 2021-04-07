package hackerrank

import "testing"

func TestAppendAndDelete(t *testing.T) {
	tests := []struct{
		desc string
		s string
		t string
		k int
		want bool
	}{
		{
			desc: "succes",
			s: "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
			t: "asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
			k: 20,
			want: true,
		},
		{
			desc: "succes",
			s: "hackerhappy",
			t: "hackerrank",
			k: 9,
			want: true,
		},
		{
			desc: "succes",
			s: "aba",
			t: "aba",
			k: 7,
			want: true,
		},
		{
			desc: "succes",
			s: "aaaaaaaaaa",
			t: "aaaaa",
			k: 7,
			want: true,
		},
		{
			desc: "failure",
			s: "qwerasdf",
			t: "qwerbsdf",
			k: 6,
			want: false,
		},
		{
			desc: "failure",
			s: "ashley",
			t: "ash",
			k: 2,
			want: false,
		},
	}

        for i, tc := range tests {
                if got := AppendAndDelete(tc.s, tc.t, tc.k); got != tc.want {
                        t.Errorf("AppendAndDelete(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
                }
        }
}
