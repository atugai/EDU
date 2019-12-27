/*

[hackerrank.com] Halloween Sale
https://www.hackerrank.com/challenges/halloween-sale

*/

package halloween_sale

// Returns the number of games possible to buy withing budget.
// p: current game price
// d: price decrement at subsequent game
// m: minimum value game can be bought
// s: total available budget
func HowManyGames(p int32, d int32, m int32, s int32) int32 {
	var (
		res int32 = 0
		sum int32 = 0
		curPrice int32 = p
	)
	for {
		// Check current sum with potential game price doesn't exceeds bugdet.
		if sum + curPrice > s {
			break
		}

		// Update game count and new sum.
		res++
		sum += curPrice

		// Update current price according to decrement and minimum game price.
		if curPrice - d > m {
			curPrice -= d
		} else {
			curPrice = m
		}
	}
	return res
}
