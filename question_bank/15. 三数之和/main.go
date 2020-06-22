package main

import (
	"fmt"
	"sort"
)

func main() {
	sum := threeSum([]int{0, 0, 0, 0})
	fmt.Println(sum)
}

func twoSum(nums []int, target int, l, r int) [][]int {
	mm := make(map[int]bool)
	var res [][]int

	for i, v := range nums[l : r+1] {
		other := target - v
		if use, ok := mm[other]; ok {
			if use {
				res = append(res, []int{-target, other, nums[i+l]})
			}
			mm[other] = false
		} else {
			mm[v] = true
		}
	}
	return res
}
func threeSum(nums []int, target int) [][]int {
	n := len(nums)
	var res [][]int

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		sum := twoSum(nums, -nums[i], i+1, n-1)
		if len(sum) == 0 {
			continue
		}
		res = append(res, sum...)
	}

	return res
}
