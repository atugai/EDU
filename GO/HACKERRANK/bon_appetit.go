/*

[hackerrank.com] Bon Appétit
https://www.hackerrank.com/challenges/bon-appetit

*/

package bon_appetit

import (
	"strconv"
)

// BonAppetit checks if fair amount was paid dugin lunch
func BonAppetit(bill []int32, k int32, b int32) string {
	var sum int32
	for _, price := range bill {
		sum += price
	}
	fair := (sum - bill[k]) / 2
	if b == fair {
		return "Bon Appetit"
	} else {
		return strconv.Itoa(int(b - fair))
	}
}
