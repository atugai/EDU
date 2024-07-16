/*

[hackerrank.com] Big Sorting
https://www.hackerrank.com/challenges/big-sorting

*/

package hackerrank

import (
	"fmt"
	"math/big"
	"sort"
)

// BigSorting returns sorted slice of big integers represented as strings.
func BigSorting(unsorted []string) []string {
	ints := []*big.Int{}

	for _, v := range unsorted {
		bi := new(big.Int)
		fmt.Sscan(v, bi)
		ints = append(ints, bi)
	}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i].Cmp(ints[j]) == -1
	})

	res := []string{}
	for _, v := range ints {
		res = append(res, v.String())
	}
	return res
}
