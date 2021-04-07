/*

[hackerrank.com] Taum and B'day
https://www.hackerrank.com/challenges/taum-and-bday

*/

package hackerrank

// TaumBDayMinimumPrice calculates minimum price for required presents
func TaumBDayMinimumPrice(b int64, w int64, bc int64, wc int64, z int64) int64 {
	var res int64
	if bc + z < wc {
		res = b*bc + w*bc + w*z
	} else if wc + z < bc {
		res = b*wc + b*z + w*wc
	} else {
		res = b*bc + w*wc
	}
	return res
}

