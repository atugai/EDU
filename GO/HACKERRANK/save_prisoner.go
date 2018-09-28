/*

[hackerrank.com] Save the Prisoner
https://www.hackerrank.com/challenges/save-the-prisoner

*/

/*

Implementation via looping here won't fit in the time limits as lagre numbers
used during tests.

*/

package save_prisoner

// SaveThePrisoner calculates for which seat bad candy would arrive
func SaveThePrisoner(pris int32, sweets int32, start int32) int32 {
	sum := start + sweets - 1
	if sum % pris == 0 {
		return pris
	}
	return sum % pris
}
