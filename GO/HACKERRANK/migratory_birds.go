/*

[hackerrank.com] Migratory Birds
https://www.hackerrank.com/challenges/migratory-birds

*/

package hackerrank

// MigratoryBirds returns most repeated element in array
// In case of sevaral elements, return lowest id
func MigratoryBirds(arr []int32) int32 {
	store := map[int32]int32{}
	var maxID, max, maxIDOld int32
	for _, n := range arr {
		if _, ok := store[n]; !ok {
			store[n] = 1
		} else {
			store[n]++
		}

		if store[n] >= max {
			maxIDOld = maxID
			maxID = n
			max = store[n]
		}
	}

	// pick lowest ID in case same number in array
	if store[maxID] == store[maxIDOld] {
		if maxID > maxIDOld {
			return maxIDOld
		}
	}
	return maxID
}
