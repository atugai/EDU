/*

[hackerrank.com] Cavity Map
https://www.hackerrank.com/challenges/cavity-map

*/

package cavity_map

// Returns updated grid map with deepest cells marked as X. Each cell of the
// input map has a value denoting its depth. Cell of the map a cavity if and
// only if this cell is not on the border of the map and each cell adjacent to
// it has strictly smaller depth.
func CavityMap(grid []string) []string {
	inMap := map[position]int32{}
	// Build current grid as map for ease of compare values.
	for i, row := range grid {
		for j, val := range row {
			inMap[position{i,j}] = val
		}
	}

	var res []string
	for i, v := range grid {
		a := ""
		// Skip top and bottom row from checking for deepest cells.
		if i == 0 || i == len(grid) - 1 {
			a = v
			res = append(res, a)
			continue
		}
		for j, val := range v {
			// Skip 1st and last column from checking for deepest cells.
			if j == 0 || j == len(v) - 1 {
				a += string(val)
				continue
			}
			// If current cell is deepest mark it as X.
			if isHigher(inMap, position{i,j}, val) {
				a += "X"
			} else {
				a += string(val)
			}
		}
		res = append(res, a)
	}
	return res
}

// Helper struct to identify cell position.
type position struct {
	row int
	col int
}

// isHigher returns true if all adjacent cells are less then current one.
func isHigher(grid map[position]int32, pos position, val int32) bool {
	// define all adjacent positions relatively to current one.
	for _, i := range []position{{-1,0}, {1,0}, {0, -1}, {0, 1}} {
		if grid[position{i.row + pos.row, i.col + pos.col}] >= val {
			return false
		}
	}
	return true
}
