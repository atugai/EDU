/*

[hackerrank.com] Picking Numbers
https://www.hackerrank.com/challenges/picking-numbers

*/

/*

This task becomes complicated if you work with raw list (time O(n^2))
Trick is to sort list before going over it. It would add extra time O(nlogn)
but this would reduce complexitiy to O(n) while looking for numbers diff.

*/

package hackerrank

import (
	"sort"
)

// PickingNumbers finds longest sequens of numbers with diff between 0 to 1
func PickingNumbers(a []int32) int32 {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	var (
		i, j int = 0, 0
		res int32 = 1
	)
	for ; i < len(a); i = j {
		var temp int32 = 1
		for j = i + 1; j < len(a); j++ {
			d := a[j] - a[i]
			if d >= 0 && d <= 1 {
				temp++
				continue
			}
			break
		}
		if temp > res {
		    res = temp
		}
	}
	return res
}
