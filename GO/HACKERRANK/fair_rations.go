/*

[hackerrank.com] Fair Rations
https://www.hackerrank.com/challenges/fair-rations

*/

package hackerrank

// FairRations returns total number of rations required to make all allocations
// even. If even allocation is not possible returs -1.
func FairRations(B []int32) int32 {
	// Numbers should be increased only between odd digit pairs, find all odds.
	var pos []int
	for i, v := range B {
		if v % 2 != 0 {
			pos = append(pos, i)
		}
	}

	// If number of odd numbers is not even then we can't allocate even number of
	// rations.
	if len(pos) % 2 != 0 {
		return -1
	}

	// Walk through all odd pairs and increase reations allocation.
	var res int32 = 0
	for i := 0; i < len(pos); i += 2 {
		beg := pos[i]
		end := pos[i+1]
		for j := beg; j < end; j++ {
			res += 2
		}
	}
	return res
}
