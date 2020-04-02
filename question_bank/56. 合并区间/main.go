package main

import "fmt"

func main() {
	ints := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println(ints)
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return res
	}

	//left:= intervals[0][0]

	left := 0
	right := 0
	var lastInt []int
	for i := range intervals[:len(intervals)-1] {

		if intervals[i][1] < intervals[i+1][0] {
			lastInt = []int{left, right}
			res = append(res, append(lastInt))
			left = intervals[i][0]
			right = intervals[i+1][1]
		} else {
			left = intervals[i][0]
			right = intervals[i+1][1]
		}
	}

	return res
}
