/*
  
[hackerrank.com] Repeated String
https://www.hackerrank.com/challenges/repeated-string/problem

*/

package repeated_string

import (
	"strings"
)

// RepeatedString calculates number of "a" in fisrt N runes of infinite string made of repeated input S.
func RepeatedString(s string, n int64) int64 {
	if int64(len(s)) >= n {
		return int64(strings.Count(s[:n], "a"))
	}
	
	fractCount := int64(strings.Count(s, "a"))
	fullFrac := n / int64(len(s))
	restFrac := n % int64(len(s))
	
	if fullFrac == 0 {
		return fractCount
	}
	
	var res int64 = fractCount * fullFrac
	res += int64(strings.Count(s[:restFrac], "a"))
	
	return res
} 
