/*

Permutations without repetitions

*/

package permutations_without_repetitions

// Permutations returns all possible permutations without repetiions of input
// elements with specified length. [1, 2] and [2, 1] considered different
// outcomes.
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
			if exists(perm, el) {
				continue
			}
			// Fix pointer problem of perm slice from outside loop: values have to be
			// copied instead re-using slice pointer.
			p := append([]int(nil), perm...)
			ret = append(ret, append(p, el))
		}
	}
	return recursivePermutations(elements, ret, length - 1)
}

func exists(l []int, el int) bool {
	for _, v := range l {
		if el == v {
			return true
		}
	}
	return false
}
