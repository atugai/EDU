package hackerrank

import (
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {
	tests := []struct{
		desc string
		word string
		h []int32
		want int32
	}{
		{
			"success",
			"abc",
			[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			9,
		},
		{
			"success",
			"zaba",
			[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
			28,
		},
	}

	for i, tc := range tests {
		if got := DesignerPdfViewer(tc.h, tc.word); got != tc.want {
			t.Errorf("TestDesignerPdfViewer(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
