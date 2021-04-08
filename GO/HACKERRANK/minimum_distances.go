/*

[hackerrank.com] Minimum Distances
https://www.hackerrank.com/challenges/save-the-prisoner

*/

package hackerrank

import "math"

// MinimumDistance returns the minimum distance between any pair of equal
// elements in the slice.
func MinimumDistance(a []int32) int32 {
	// Build helper dictionary to store all positions for each slice item.
	nums := map[int32][]int{}
	for pos, val := range a {
		ar, ok := nums[val]
		if ok {
			ar = append(ar, pos)
			nums[val] = ar
		} else {
			nums[val] = []int{pos}
		}
	}

	// Initialise maximum distance for input slice lenght.
	var res int32 = int32(len(a))
	for _, positions := range nums {
		// Ignore if similar elements are not 2.
		if len(positions) != 2 {
			continue
		}
		delta := int32(math.Abs(float64(positions[0] - positions[1])))
		if delta < res {
			res = delta
		}
	}

	// Check if any distances were found.
	if res == int32(len(a)) {
		return -1
	}
	return res
}
