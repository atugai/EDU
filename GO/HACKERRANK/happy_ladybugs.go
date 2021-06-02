/*

  [hackerrank.com] Happy Ladybugs
  https://www.hackerrank.com/challenges/happy-ladybugs

*/

package hackerrank

// import "fmt"

// isEmpty returns true if rune is empty char "_".
func isEmpty(in rune) bool {
	return in == 95
}

// hasEmpty returns true if input has at least one empty char "_".
func hasEmpty(in map[rune]int) bool {
	if _, ok := in[95]; ok {
		return true
	}
	return false
}

// allEmpty returns true if input has only empty chars "_".
func allEmpty(in map[rune]int) bool {
	if len(in) == 1 && hasEmpty(in) {
		return true
	}
	return false
}

// hasMinimumBugs returns true if input has at least 2 runes of same char,
// minimum number to make seq happy for each symbol except empty char "_".
func hasMinimumBugs(in map[rune]int) bool {
	for k, v := range in {
		if v < 2 && !isEmpty(k) {
			return false
		}
	}
	return true
}

// isHappy returns true if sequense already happy - checks that at least 2
// adjacent symbols are same.
func isHappy(in string) bool {
	size := len(in)
	if size < 2 {
		return false
	}
	for i := 1; i < size-1; i++ {
		if in[i] != in[i-1] && in[i] != in[i+1] {
			return false
		}
	}
	if in[0] != in[1] || in[size-1] != in[size-2] {
		return false
	}
	return true
}

// HappyLadybugs returns true if input is a happy sequence or can be made happy.
func HappyLadybugs(b string) bool {
	base := map[rune]int{}
	for _, r := range b {
		base[r] += 1
	}
	if len(base) == 0 || allEmpty(base) {
		return true
	}
	if hasMinimumBugs(base) && hasEmpty(base) {
		return true
	}
	if isHappy(b) {
		return true
	}
	return false
}
