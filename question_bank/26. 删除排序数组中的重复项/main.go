package main

import "fmt"

func main() {
	duplicates := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(duplicates)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var pre = nums[0]
	var j = 1
	for i := j; i < len(nums); i++ {
		if nums[i] != pre {
			nums[j] = nums[i]
			pre = nums[i]
			j++
		}
	}
	fmt.Println(j)
	fmt.Println(nums)
	return j
}
