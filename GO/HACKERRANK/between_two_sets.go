/*

[hackerrank.com] Between Two Sets
https://www.hackerrank.com/challenges/between-two-sets

You have two arrays of integers. Determine all integers that satisfy the following two conditions:
* The elements of the first array are all factors of the integer being considered (slice a)
* The integer being considered is a factor of all elements of the second array (slice b)

These numbers are referred to as being between the two arrays.
You must determine how many such numbers exist.

*/

package hackerrank

// factorsOf checks if all numbers in list are factors of t 
func factorsOf(ls []int32, t int32) bool {
    for _, n := range ls {
        if t % n != 0 {
            return false
        }
    }
    return true
}

// isFactor checks if t is factor of all numbers in list
func isFactor(ls []int32, t int32) bool {
    for _, n := range ls {
        if n % t != 0 {
            return false
        }
    }
    return true
}

// GetTotal checks total number of integers between 2 slices
func GetTotal(a []int32, b []int32) int32 {
    var res int32 = 0
    for i := a[len(a)-1]; i <=  b[0]; i++ {
        if factorsOf(a, i) && isFactor(b, i) {
            res++
        }
    }
    return res
}
