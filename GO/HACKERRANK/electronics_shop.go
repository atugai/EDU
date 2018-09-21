/*

[hackerrank.com] Electronics Shop
https://www.hackerrank.com/challenges/electronics-shop

*/

package electronics_shop

// GetMoneySpent calculates maximum spendable budget for kb and drive
func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var (
		res int32
		found bool
	)
	for _, kb := range keyboards {
		for _, dr := range drives {
			r := kb + dr
			if r <= b && r >= res {
				found = true
				res = r
			}
		}
	}
	if found {
		return res
	}
	return -1
}
