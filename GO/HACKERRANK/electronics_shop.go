/*

[hackerrank.com] Electronics Shop
https://www.hackerrank.com/challenges/electronics-shop

*/

package electronics_shop

// GetMoneySpent calculates maximum spendable budget for kb and drive
func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var (
		outcomes []int32
		res int32
		found bool
	)
	for _, kb := range keyboards {
		for _, dr := range drives {
			outcomes = append(outcomes, kb + dr)
		}
	}
	for _, r := range outcomes {
		if r <= b && r >= res {
			found = true
			res = r
		}
	}
	if found {
		return res
	}
	return -1
}
