/*
  
[hackerrank.com] Jumping on the Clouds
https://www.hackerrank.com/challenges/jumping-on-the-clouds

*/

package jumping_clouds_basic

// JumpingOnClouds jumps over clouds and calculates minimum number of steps to get the end avoiding thunderclouds.
func JumpingOnClouds(c []int32) int32 {
	var res int32 = 0
	for i := 0; i <= len(c) - 2; {
		if (i + 2) < len(c) {
			if c[i + 2] == 1 {
				i = i + 1
			} else {
				i = i + 2
			}
		} else {
			i = i + 1
		}
		res += 1
	}
	return res
}
