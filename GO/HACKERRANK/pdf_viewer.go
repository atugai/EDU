/*

[hackerrank.com] Designer PDF Viewer
https://www.hackerrank.com/challenges/designer-pdf-viewer

*/

package pdf_viewer

// DesignerPdfViewer calculates highligt size for provided words
func DesignerPdfViewer(h []int32, word string) int32 {
	heights := map[int32]int32{}
	for i, v := range h {
		heights[rune('a') + int32(i)] = v
	}

	var max int32 = 0
	for _, v := range word {
		if  heights[v] > max {
			max = heights[v]
		}
	}
	return max * int32(len(word))
}
