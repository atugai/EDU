/*
  
[hackerrank.com] Repeated String
https://www.hackerrank.com/challenges/repeated-string/problem

*/

package repeated_string

import (
	"log"
	"bytes"
)

// RepeatedString calculates number of "a" in fisrt N runes of infinite string made of repeated input S.
func RepeatedString(s string, n int64) int64 {
	if len(s) == 1 && s == "a" {
		return n
	}
	
	var bs bytes.Buffer
	for int64(bs.Len()) < n {
		bs.WriteString(s)
		log.Println("COMBINING string")
	}
	res := bs.String()
	
	var sum int64 = 0
	for i := 0; int64(i) < n; i++ {
		log.Println("Searching for 'a'")
		if res[i] == 'a' {
			sum += 1
		}
	}
	return sum  
} 
