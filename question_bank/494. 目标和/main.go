package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 2, -2}
	//1 : 1,1,1,1
	//
	ways := findTargetSumWays(nums, 3)
	fmt.Println(ways)
}

func findTargetSumWays(nums []int, S int) int {
	var res int
	resMap := make(map[int]map[int]int)
	res = findHelper(nums, 0, 0, S, resMap)
	return res
}

func findHelper(nums []int, pos int, sum int, S int, resMap map[int]map[int]int) int {
	if sum == S && pos == len(nums) {
		return 1
	} else if pos == len(nums) && sum != S {
		return 0
	}

	if _, ok := resMap[pos]; ok {
		if v, ok := resMap[pos][sum]; ok {
			return v
		}
	}

	left := findHelper(nums, pos+1, sum+nums[pos], S, resMap)
	right := findHelper(nums, pos+1, sum-nums[pos], S, resMap)
	if _, ok := resMap[pos]; !ok {
		resMap[pos] = make(map[int]int)
	}
	resMap[pos][sum] = left + right
	return left + right
}
