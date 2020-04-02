package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	islands := numIslands(grid)
	fmt.Println(islands)
}

// i , j
var directed = [4][2]int{
	{0, 1},  //向右
	{1, 0},  //向下
	{0, -1}, //向左
	{-1, 0}, //向上
}

func numIslands(grid [][]byte) int {
	var res int
	for col, v := range grid {
		for row, v := range v {
			if v == '1' {
				grid[col][row] = '-'
				res++
				numIslandsHelper(grid, col, row)
			}
		}
	}
	return res
}

func numIslandsHelper(grid [][]byte, col, row int) {

	for _, v := range directed {
		i, j := col+v[0], row+v[1]
		if (i >= 0 && i < len(grid)) && (j >= 0 && j < len(grid[0])) && (grid[i][j] == '1') {
			grid[i][j] = '-'
			numIslandsHelper(grid, i, j)
		}
	}
}
