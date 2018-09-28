/*

[hackerrank.com] Beautiful Days at the Movies
https://www.hackerrank.com/challenges/beautiful-days-at-the-movies

*/

package beautiful_days

import (
	"math"
	"strconv"
)

// reverse helper function to reverse integer
func reverse(in int32) (int32, error) {
	num := strconv.Itoa(int(in))

	var rev string
	for _, s := range num {
		rev = string(s) + rev
	}
	res, err := strconv.Atoi(rev)
	return int32(res), err
}

// BeautifulDays calculate number of days in specified range
func BeautifulDays(i int32, j int32, k int32) int32 {
	var res, n int32

	for n = i; n <= j; n++ {
		nRev, _ := reverse(n)
		if int32(math.Abs(float64(n - nRev))) % k == 0 {
			res++
		}
	}
	return res
}
