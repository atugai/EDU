/*

[hackerrank.com] Counting Valleys
https://www.hackerrank.com/challenges/counting-valleys

*/

package counting_valleys

// CountingValleys counts number of valleys in journey string path
func CountingValleys(s string) int32 {
	var (
		res, position int32
		prevUp bool
	)
	for _, r := range []rune(s) {
		if r == rune('U') {
			position +=1
			prevUp = true
		} else {
			position -= 1
			prevUp = false
		}
		if position == 0 && prevUp {
			res++
		}
	}
	return res
}
