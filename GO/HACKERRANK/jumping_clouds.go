/*
  
[hackerrank.com] Jumping on the Clouds Revisited
https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited

*/

package hackerrank

// JumpingOnClouds jumps over clouds and calculates resulting energy afterwards.
func JumpingOnClouds(c []int32, k int32) int32 {
	var ( 
		e int32 = 100
		cnt int32 = 0
	)
	for cnt < int32(len(c)) {
		cnt += k
		e -= 1
		if c[cnt % int32(len(c))] == 1 {
			e -= 2
		}
	}
	return e
}
