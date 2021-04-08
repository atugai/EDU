/*

[hackerrank.com] Manasa and Stones
https://www.hackerrank.com/challenges/manasa-and-stones

*/

package hackerrank

import (
	"fmt"
	"sort"
)

// Generator helper struct to generate permutations. This are not real
// permutations - since we don't need to know all outomes, Ex. [1, 1, 2],
// [1, 2, 1], [2, 1, 1] would be same outcomes as we would use them only to get
// sum value of elements, So generator would return only [1, 1, 2], [1, 2, 2]
// and [2, 2, 2] outcomes + base outcome [1, 1, 1].
type Generator struct {
	// Pos current postiion to update in resulting outcome.
	pos int
	// Val value to replace original one in produced outcome.
	val int32
	// Current generated outcome.
	current []int32
}

// NewGenerator pre-build generator and initializes base outcome.
func NewGenerator(l, a, b int32) *Generator {
	// Build base outcome.
	v := make([]int32, l)
	for i := 0; i < int(l); i++ {
		v[i] = a
	}

	return &Generator{
		val:     b,
		current: v,
	}
}

// Base returns base/initial outcome.
func (g *Generator) Base() []int32 {
	return g.current
}

// Next returns next generated outcome. Returns out of range error when no more
// outcomes to be produced.
func (g *Generator) Next() ([]int32, error) {
	if g.pos >= len(g.current) {
		return nil, fmt.Errorf("out of range")
	}
	// Update outcome with next element and increase current position.
	g.current[g.pos] = g.val
	g.pos++
	return g.current, nil
}

// Stones computes all possible stone values at the end of the road. Takes as
// input number of stones on the trail and 2 possible distances between stones.
func Stones(n int32, a int32, b int32) []int32 {
	// Initialise generator and get base outcome.
	g := NewGenerator(n-1, a, b)
	base := compute(g.Base())

	// Create temp map and insert base outcome.
	tempMap := map[int32]int32{}
	tempMap[base[len(base)-1]] = -1

	// fetch all outcomes from generator and insert in map.
	res := []int32{}
	for {
		perm, err := g.Next()
		if err != nil {
			break
		}
		roadStones := compute(perm)
		lastStone := roadStones[len(roadStones)-1]

		if _, ok := tempMap[lastStone]; !ok {
			tempMap[lastStone] = -1
		}
	}

	// Convert calculated stones into slice and sort before return result.
	for k := range tempMap {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

// Compute base on permutations sequence with stone vaules differences calculate
// original sequence of sotnes.
func compute(in []int32) []int32 {
	res := []int32{0}
	var s int32
	for _, v := range in {
		s += v
		res = append(res, s)
	}
	return res
}
