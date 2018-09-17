/*

[hackerrank.com] Migratory Birds
https://www.hackerrank.com/challenges/migratory-birds

*/

package migratory_birds

// MigratoryBirds returns most repeated element in array
// In case of sevaral elements, return lowest id
func MigratoryBirds(arr []int32) int32 {
	store := map[int32]int32{}
	var maxId, max, maxIdOld int32
	for _, n := range arr {
		if _, ok := store[n]; !ok {
			store[n] = 1
		} else {
			store[n] += 1
		}

		if store[n] >= max {
			maxIdOld = maxId
			maxId = n
			max = store[n]
		}
	}

	// pick lowest ID in case same number in array
	if store[maxId] == store[maxIdOld] {
		if maxId > maxIdOld {
			return maxIdOld
		}
	}
	return maxId
}
