/*
  
[hackerrank.com] Sequence Equation
https://www.hackerrank.com/challenges/permutation-equation

*/

package hackerrank

// SequenceEquation returns slice of integers pointed by integers from 0 to n in input slice.
func SequenceEquation(p []int32) []int32 {
	valToPos := map[int32]int32{}
	for pos, val := range p {
		valToPos[val] = int32(pos) + 1
	}

	var res []int32
	for i := 1; i <= len(p); i++ {
		res = append(res, valToPos[valToPos[int32(i)]])
	}
	return res
}
