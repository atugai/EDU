/*

[hackerrank.com] Circular Array Rotation
https://www.hackerrank.com/challenges/circular-array-rotation

*/

/*

This implementation is not mamory efficient as requires to store modified slice
TODO: optimize

*/

package hackerrank

// CircularArrayRotation calculates new array element positions after right circullar rotation
func CircularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	mod := make([]int32, len(a))
		for i, v := range a {
			mod[(i + int(k)) % len(a)] = v
		}

	res := make([]int32, len(queries))
	for i, v := range queries {
		res[i] = mod[v]
	}
	return res
}
