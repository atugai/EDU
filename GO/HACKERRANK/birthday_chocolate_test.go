package hackerrank

import (
	"testing"
)

func TestBirthday(t *testing.T) {
	tests := []struct{
		desc string
		input []int32
		day, month, want int32
	}{
		{
			desc: "non-zero result",
			input: []int32{1, 2, 1, 3, 2},
			day: 3,
			month: 2,
			want: 2,
		},
		{
			desc: "zero input",
			input: []int32{1, 1, 1, 1, 1, 1},
			day: 3,
			month: 2,
			want: 0,
		},
		{
			desc: "non-zero result",
			input: []int32{4},
			day: 4,
			month: 1,
			want: 1,
		},
	}

	for i, tc := range tests {
		if got := Birthday(tc.input, tc.day, tc.month); got != tc.want {
			t.Errorf("TestBirthday(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
