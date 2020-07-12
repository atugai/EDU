/*

Permutations with repetitions

*/

package permutations_with_repetitions

// Improved version of Permutations. This approach doesn't retun whole list at
// once but rather uses generator and function Next() to get next permutation.
// This approach doesn't use recursion at all and has simple way to generate
// outcomes.
// If we look at standard permutation outcome it consists of fixed number input
// elements: [e1, e2, e3, ... en]. This means we know number of elements in
// resulting outcome and we know the elements. Only what is left is to generate
// sequense order for theese elements. This approach is used in generator - at
// each call Next() we just generate sequence of number - positions of input
// elements in original list.

import (
	"errors"
	"math"
)

// Generator holder struct to store all required counters.
type Generator struct {
	// Elements items from which permutation is built.
	elements []int
	// Total number of  outcomes possible for required length and input elements.
	total int
	// Current outcome sequence number.
	current int
	// Position current outcome positions of elements in original list.
	position []int
}

// NewGenerator builds permutations generator.
func NewGenerator(length int, elements []int) *Generator {
	// Calculate number of possible outcomes.
	t := int(math.Pow(float64(len(elements)), float64(length)))

	return &Generator{
		elements: elements,
		total: t,
		position: make([]int, length),
	}
}

// Next returns next outcome. Returns out of range error when no more outcomes
// to be generated.
func (g *Generator) Next() ([]int, error) {
	// Check  more outcomes can be generated.
	if g.current >= g.total {
		return nil, errors.New("out of range")
	}

	// Build current outcome. Since we update sequence numbers from left, outcome
	// generated from rifht, so build result starting from end of slice.
	res := []int{}
	for i := len(g.position) - 1; i >= 0; i-- {
		res = append(res, g.elements[g.position[i]])
	}

	// Update position numbers and increase current number.
	g.updateCounter()

	return res, nil
}

// updateCounter updates cuurent outcome number and current outcome position
// numbers.
func (g *Generator) updateCounter() {
	// Loop over each position and increase it by 1. If position already max
	// value, then drop it to 0 and increase next one by 1.
	for i, _ := range g.position {
		g.position[i]++
		if g.position[i] < int(len(g.elements)) {
			break
		}
		g.position[i] = 0
	}
	g.current++
}
