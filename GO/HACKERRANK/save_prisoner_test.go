package save_prisoner

import (
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {
	tests := []struct{
		desc string
		numPris int32
		numSweets int32
		startPos int32
		want int32
	}{
		{
			"success",
			352926151,
			380324688,
			94730870,
			122129406,
		},
		{
			"success",
			7,
			19,
			2,
			6,
		},
		{
			"success",
			5,
			2,
			2,
			3,
		},
	}

	for i, tc := range tests {
		if got := SaveThePrisoner(tc.numPris, tc.numSweets, tc.startPos); got != tc.want {
			t.Errorf("SaveThePrisoner(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
