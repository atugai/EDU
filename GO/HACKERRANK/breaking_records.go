/*

[hackerrank.com] Breaking the Records
https://www.hackerrank.com/challenges/breaking-best-and-worst-records

*/

package hackerrank

// BreakingRecords calculates numer of times Max and Min were exceeded
func BreakingRecords(scores []int32) []int32 {
	var minBr, maxBr, min, max int32
	min = scores[0]
	max = scores[0]
	for _, s := range scores {
		if s > max {
			max = s
			maxBr++
		}
		if s < min {
			min = s
			minBr++
		}
	}
	return []int32{maxBr, minBr}
}
