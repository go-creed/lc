package main

import (
	"fmt"
	"sort"
)

func main() {
	window := medianSlidingWindow([]int{1, 4, 2, 3}, 4)
	fmt.Println(window)
}

//
func medianSlidingWindow(nums []int, k int) []float64 {
	l := len(nums)
	var res []float64
	if l == 0 {
		return res
	}

	var right int
	var left int
	var isFull bool
	var tf bool
	if k&1 == 0 {
		tf = true
	}

	for x := 0; x <= l; x++ {
		if !isFull {
			right++
			if right >= k {
				isFull = true
			}
			continue
		}
		ints := append([]int{}, nums[left:right]...)
		sort.Ints(ints)
		fmt.Println(ints)
		pos := (k) / 2
		var tmp float64
		if tf {
			tmp = float64(ints[pos]+ints[pos-1]) / 2
		} else {
			tmp = float64(ints[pos])
		}
		left++
		right++
		res = append(res, tmp)
	}

	return res
}
