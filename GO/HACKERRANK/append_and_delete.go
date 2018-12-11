/*
  
[hackerrank.com] Append and Delete
https://www.hackerrank.com/challenges/append-and-delete

*/

package append_and_delete

// AppendAndDelete checks if possible to get string t from string s for k number of steps/changes.
func AppendAndDelete(s string, t string, k int) bool {
	// Check if we can delete all letters within limit k.
	if len(s) + len(t) <= k {
		return true
	}

	// calculate how much of letters are same in 2 words.
	i := 0
	for i < min(len(s), len(t)) {
		if s[i] != t[i] {
			break
		}
		i += 1
	}

	// calculate  how many steps we need to make words equal.
	delta := len(s) - i + len(t) - i

	// If diff letters either empty or equels to number of steps we have, we're good.
	if delta == k || delta == 0 {
		return true
	}

	// if more changes to be made then we have steps in k we're not good.
	if delta > k {
		return false
	}

	// Check if nuber of steps left can be split equally between 2 words to make them equal.
	if ((k - delta) % 2 == 0) {
		return true
	}
	return false
}

// min helper function to get min out of 2 ints.
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
