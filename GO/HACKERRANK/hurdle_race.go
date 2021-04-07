/*

[hackerrank.com] The Hurdle Race
https://www.hackerrank.com/challenges/the-hurdle-race

*/

package hackerrank

// HurdleRace calculates number of potions to take to jump over hieghts
func HurdleRace(k int32, height []int32) int32 {
	var max int32 = 0
	for _, v := range height {
		if v > max {
			max = v
		}
	}
	if max > k {
		return (max - k)
	}
	return 0
}
