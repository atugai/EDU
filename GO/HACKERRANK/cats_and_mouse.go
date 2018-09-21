/*

[hackerrank.com] Cats and a Mouse
https://www.hackerrank.com/challenges/cats-and-a-mouse

*/

package cats_and_mouse

import (
	"math"
)

// CatAndMouse checks how far each cat from mouse
func CatAndMouse(x int32, y int32, z int32) string {
	catA := math.Abs(float64(z - x))
	catB := math.Abs(float64(z - y))
		switch {
		case catA > catB:
			return "Cat B"
		case catB > catA:
			return "Cat A"
		default:
			return "Mouse C"
	}
}
