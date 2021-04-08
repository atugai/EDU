/*

[hackerrank.com] Flatland Space Stations
https://www.hackerrank.com/challenges/flatland-space-stations

*/

package hackerrank

import "math"

// FlatlandSpaceStations returns the maximum distance required to get from city
// to space station.
// n: number of cities.
// c: slice describing which city has space station.
// Example: [0, 4] means city 0 and 4 has space station.
func FlatlandSpaceStations(n int32, c []int32) int32 {
	var res int32 = math.MinInt32
	for i := 0; i < int(n); i++ {
		var dst int32 = math.MaxInt32
		for _, v := range c {
			if val := int32(math.Abs(float64(v - int32(i)))); val < dst {
				dst = val
			}
		}
		if res < dst {
			res = dst
		}
	}
	return res
}
