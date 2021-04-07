/*

[hackerrank.com] Drawing Book
https://www.hackerrank.com/challenges/drawing-book

*/

/*

Challenge here is to find way how to represent number of turns from front and
from back of the book:

Turns F | Pages | Turns B
=========================
      0 |   ' 1 | 4
      1 | 2 ' 3 | 3
      2 | 4 ' 5 | 2
      3 | 6 ' 7 | 1
      4 | 8 '   | 0

*/

package hackerrank

// turnFront helper func to find number of turns from front
func turnFront(p int32) int32 {
	if p % 2 == 0 {
		return p/2
	} else {
		return (p - 1)/2
	}
}

// turnBack helper func to find number of turns from back
func turnBack(n int32, p int32) int32 {
	if p % 2 == 0 {
		return (n - p)/2
	} else {
		return (n - p + 1)/2
	}
}

// PageCount returns minimum number of page turns to get required page
func PageCount(n int32, p int32) int32 {
	if turnFront(p) > turnBack(n, p) {
		return turnBack(n, p)
	}
	return turnFront(p)
}
