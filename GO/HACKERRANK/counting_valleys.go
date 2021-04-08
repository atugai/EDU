/*

[hackerrank.com] Counting Valleys
https://www.hackerrank.com/challenges/counting-valleys

*/

package hackerrank

// CountingValleys counts number of valleys in journey string path
func CountingValleys(s string) int32 {
	var (
		res, position int32
		prevUp bool
	)
	for _, r := range []rune(s) {
		if r == rune('U') {
			position++
			prevUp = true
		} else {
			position--
			prevUp = false
		}
		if position == 0 && prevUp {
			res++
		}
	}
	return res
}
