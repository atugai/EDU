/*

Permutations with repetitions

*/

package permutations_with_repetitions

// WARNING !!! Brute force implementation, fails on large outcome lengths.

// Permutations returns all possible permutations with repetitions of input
// elements with specified length. Repeated elements considered as different.
// Elements order matters so [1, 2] and [2, 1] are different outcomes.
// Allows to get base elements as outcomes and intermediate outcomes with
// length less then specified by setting full to true.
func Permutations(elements []int, length int, full bool) [][]int {
	base := [][]int{}
	// Check if permutation length is not 0.
	if length <= 0 {
		return base
	}

	// Prepare base permutation with length 1.
	for _, el := range elements {
		base = append(base, []int{el})
	}

	return recursivePermutations(elements, base, length - 1, full)
}

// recursivePermutations recursively builds all possible permutations from
// elements and pre-build results from previous iterations.
func recursivePermutations(elements []int, res [][]int, length int, full bool) [][]int {
	// Return results immediately if we reach length of permutation sequence.
	if length == 0 {
		return res
	}

	// If full permutations required add empty slice to insert base values.
	if full {
		res = append([][]int{{}}, res...)
	}

	// Build current iteration permutations and call next iteration.
	ret := [][]int{}
	for _, perm := range res {
		for _, el := range elements {
			// Fix pointer problem of perm slice from outside loop: values have to be
			// copied instead re-using slice pointer.
			p := append([]int(nil), perm...)
			ret = append(ret, append(p, el))
		}
	}
	return recursivePermutations(elements, ret, length - 1, full)
}
