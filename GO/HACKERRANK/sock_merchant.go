/*

[hackerrank.com] Sock Merchant
https://www.hackerrank.com/challenges/sock-merchant

*/

package sock_merchant

// SockMerchant finds number of sock pairs in array of socks
func SockMerchant(ar []int32) int32 {
	socks := map[int32]int32{}
	var res int32
	for _, s := range ar {
		_, ok := socks[s]
		if !ok {
			socks[s] = 1
		} else {
			socks[s]++
			if socks[s] % 2 == 0 {
				res++
			}
		}
	}
	return res
}
