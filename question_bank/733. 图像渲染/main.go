package main

import "fmt"

func main() {
	var iamge [][]int = [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	fill := floodFill(iamge, 1, 1, 2)

	for _, ints := range fill {
		for _, i2 := range ints {
			fmt.Print(i2, " ")
		}
		fmt.Println()
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var color int = image[sr][sc]
	if color != newColor {
		dfs(image, sr, sc, color, newColor)
	}
	return image
}

func dfs(image [][]int, r, c, color, newColor int) {
	if image[r][c] != color {
		return
	}
	image[r][c] = newColor
	if r >= 1 {
		dfs(image, r-1, c, color, newColor)
	}
	if c >= 1 {
		dfs(image, r, c-1, color, newColor)
	}
	if r+1 < len(image) {
		dfs(image, r+1, c, color, newColor)
	}
	if c+1 < len(image[0]) {
		dfs(image, r, c+1, color, newColor)
	}
}
