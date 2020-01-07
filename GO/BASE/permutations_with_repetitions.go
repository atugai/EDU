/*

Permutations with repetitions

*/

package permutations_with_repetitions

// Permutations returns all possible permutations with repetiions of input
// elements with specified length. Repeated elements considered as different.
func Permutations(elements []int, length int) [][]int {
	base := [][]int{}
	// Check if permutation length is not 0.
	if length <= 0 {
		return base
	}

	// Prepare base permutation with length 1.
	for _, el := range elements {
		base = append(base, []int{el})
	}

	// Start recursive permutation builds with decremented length.
	return recursivePermutations(elements, base, length - 1)
}

// recursivePermutations recursively builds all possible permutations from
// elements and pre-build results from previous iterations.
func recursivePermutations(elements []int, res [][]int, length int) [][]int {
	// Return results immediately if we reach length of permutation sequence.
	if length == 0 {
		return res
	}

	// Build current iteration permutations and call next iteration.
	ret := [][]int{}
	for _, perm := range res {
		for _, el := range elements {
			ret = append(ret, append(perm, el))
		}
	}
	return recursivePermutations(elements, ret, length - 1)
}
