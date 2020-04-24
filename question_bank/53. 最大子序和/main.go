package main

import (
	"fmt"
	"math"
)

func main() {
	array := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(array)
}

func maxSubArray(nums []int) int {
	var maxSum int = nums[0]
	var curSum int = maxSum
	for i := 1; i < len(nums); i++ {
		curSum = int(math.Max(float64(nums[i]+curSum), float64(nums[i])))
		maxSum = int(math.Max(float64(curSum), float64(maxSum)))
	}
	return maxSum
}
