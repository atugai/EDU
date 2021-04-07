/*

[hackerrank.com] Viral Advertising
https://www.hackerrank.com/challenges/strange-advertising

*/

package hackerrank

// ViralAdvertising calculates number of likes by specified day
func ViralAdvertising(n int32) int32 {
	var shared int32 = 5
	var likes, res int32

	for i := 0; i < int(n); i++ {
		likes = int32(shared/2)
		res += likes
		shared = likes * 3
	}
	return res
}
