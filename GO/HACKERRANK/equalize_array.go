/*

[hackerrank.com] Equalize the Array
https://www.hackerrank.com/challenges/equality-in-a-array

*/

package hackerrank

import (
	"sort"
)

// EqualizeArray calculates minimum number of lelements to remove from arra to leave only equal elements.
func EqualizeArray(arr []int32) int32 {
	els := map[int32]int{}
	for _, v := range arr {
		if _, ok := els[v]; !ok {
			els[v] = 1
		} else {
			els[v]++
		}
	}

	nums := []int{}
	for _, v := range els {
		nums = append(nums, v)
	}
	sort.Ints(nums)

	var res int32 = 0
	for i := 0; i < len(nums) - 1; i++ {
		res += int32(nums[i])
	}
	return res
}
