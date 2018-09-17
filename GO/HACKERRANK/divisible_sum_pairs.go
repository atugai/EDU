/*

[hackerrank.com] Divisible Sum Pairs
https://www.hackerrank.com/challenges/divisible-sum-pairs

*/

package divisible_sum_pairs

// DivisibleSumPairs finds all sum pairs which are divisible by k
func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var res int32
	for i, _ := range ar {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i] + ar[j]) % k == 0 {
				res++
			}
		}
	}
	return res
}
