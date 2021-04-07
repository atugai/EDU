/*

[hackerrank.com] Day of the Programmer
https://www.hackerrank.com/challenges/day-of-the-programmer

*/

package hackerrank

import (
	"strconv"
)

// isLeap helper function to check if year after 1918 is leap
func isLeap(y int32) bool {
	if (y % 400 == 0) || ((y % 4 == 0) && (y % 100 != 0)) {
		return true
	}
	return false
}

// DayOfProgrammer returns Programmer day (256th day) from provided year
func DayOfProgrammer(year int32) string {
	switch {
	case (year >= 1919):
		if isLeap(year) {
			return "12.09." + strconv.Itoa(int(year))
		}
		return "13.09." + strconv.Itoa(int(year))
	case (year <= 1917):
		if (year % 4 == 0) {
			return "12.09." + strconv.Itoa(int(year))
		}
		return "13.09." + strconv.Itoa(int(year))
	}
	return "26.09.1918"
}
