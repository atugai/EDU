/*

[hackerrank.com] Birthday Chocolate
https://www.hackerrank.com/challenges/the-birthday-bar

*/

package hackerrank

// Birthday calculates chocolate pieces allocation for day and month
func Birthday(s []int32, d int32, m int32) int32 {
	var res, sum, num int32
	for i := range s {
		sum, num = 0,0
		for j := i; j < len(s); j++ {
			sum += s[j]
			num++
			if (sum == d) && (m == num) {
				res++
				break
			}
			if (sum > d) || (num > m) {
				break
			}
		}
	}
	return res
}
