package main

import (
	"fmt"
	"math"
)

func main() {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}
	wallsAndGates(rooms)

	fmt.Println(rooms)
}

func wallsAndGates(rooms [][]int) {
	for i, v := range rooms {
		for j, v := range v {
			if v == 0 {
				wallHelper(rooms, i, j, 0)
			}
		}
	}
}

func wallHelper(rooms [][]int, i, j int, distance int) int {
	m := len(rooms)
	n := len(rooms[0])
	rooms[i][j] = int(math.Min(float64(rooms[i][j]), float64(distance)))
	if i > 0 && rooms[i-1][j] > distance+1 {
		wallHelper(rooms, i-1, j, distance+1)
	}
	if j > 0 && rooms[i][j-1] > distance+1 {
		wallHelper(rooms, i, j-1, distance+1)
	}
	if i < m-1 && rooms[i+1][j] > distance+1 {
		wallHelper(rooms, i+1, j, distance+1)
	}
	if j < n-1 && rooms[i][j+1] > distance+1 {
		wallHelper(rooms, i, j+1, distance+1)
	}
	return 0
}
