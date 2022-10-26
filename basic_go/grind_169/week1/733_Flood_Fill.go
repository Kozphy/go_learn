package week1

import "fmt"

// func dfs(image *[][]int, sr int, sc int, color int, rows int, cols int, source int) {
// 	if sr < 0 || sr >= rows || sc < 0 || sc >= cols {
// 		return
// 	} else if (*image)[sr][sc] != source {
// 		return
// 	}

// 	image[sr][sc] == color

// 	dfs(image, sr-1, sc, color, rows, cols, source)
// 	dfs(image, sr+1, sc, color, rows, cols, source)
// 	dfs(image, sr, sc-1, color, rows, cols, source)
// 	dfs(image, sr, sc+1, color, rows, cols, source)
// }

func floodFill(image *[][]int, sr int, sc int, color int) *[][]int {
	// if color == (*image)[sr][sc] {
	// 	return image
	// }
	// rows := len(*image)
	// cols := len((*image)[0])
	// source := (*image)[sr][sc]
	// dfs(image, sr, sc, color, rows, cols, source)
	// return image
	return image
}

// TODO: complete this
func Execute_floodFill() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	floodFill(&image, 1, 1, 2)
	fmt.Println(image)
}
