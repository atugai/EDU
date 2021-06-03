/*

[hackerrank.com] Strange Counter
https://www.hackerrank.com/challenges/strange-code

*/

package hackerrank

// StrangeCounter returns strange counter value at specified time.
func StrangeCounter(t int64) int64 {
	start, end := int64(1), int64(3)
	timer := end
	for {
		if t <= end && t >= start {
			break
		}
		start = end + 1
		timer = timer * 2
		end = end + timer
	}
	return start + timer - t
}
