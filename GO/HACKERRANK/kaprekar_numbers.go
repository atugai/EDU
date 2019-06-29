/*
  
[hackerrank.com] Kaprekar Numbers
https://www.hackerrank.com/challenges/kaprekar-numbers

*/

package kaprekar_numbers

import (
	"strconv"
)

// KaprekarNumbers returns Kaprekar numbers in range p and q inclusive.
func KaprekarNumbers(p int32, q int32) []int32 {
	var res []int32
	for i := p; i <= q; i++ {
		if isKaprekar(i) {
			res = append(res, i)
		}
	}
	return res
}

// isKaprekar validates if input is KAprekar number
func isKaprekar(n int32) bool {
    v := int64(n) * int64(n)
    vs := strconv.FormatInt(v, 10)
    l, r := vs[len(vs)/2:], vs[0:len(vs)/2]

    lInt, _ := strconv.ParseInt(l, 10, 64)
    rInt, _ := strconv.ParseInt(r, 10, 64)

    return n == int32(lInt + rInt)
} 
