package cats_and_mouse

import (
	"testing"
)

func TestCatAndMouse(t *testing.T) {
	tests := []struct{
		desc string
		catA int32
		catB int32
		mouse int32
		want string
	}{
		{
			"cat B closer",
			1,
			2,
			3,
			"Cat B",
		},
		{
			"mouse escapes",
			1,
			3,
			2,
			"Mouse C",
		},

	}

	for i, tc := range tests {
		if got := CatAndMouse(tc.catA, tc.catB, tc.mouse); got != tc.want {
			t.Errorf("CatAndMouse(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
