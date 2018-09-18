package programmer_day

import (
	"testing"
)

func TestDayOfProgrammer(t *testing.T) {
	tests := []struct{
		desc string
		year int32
		want string
	}{
		{
			"leap year after 1918",
			2016,
			"12.09.2016",
		},
		{
			"not leap year after 1918",
			2017,
			"13.09.2017",
		},
		{
			"leap year before 1800",
			1800,
			"12.09.1800",
		},
		{
			"not leap year before 1800",
			1799,
			"13.09.1799",
		},
		{
			"year 1918",
			1918,
			"26.09.1918",
		},
	}

	for i, tc := range tests {
		if got := DayOfProgrammer(tc.year); got != tc.want {
			t.Errorf("TestDayOfProgrammer(%d): %s: unexpected result got: %v, want: %v", i, tc.desc, got, tc.want)
		}
	}
}
