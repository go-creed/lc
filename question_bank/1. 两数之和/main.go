package main

import "fmt"

func main() {
	sum := twoSum([]int{3,2,4}, 6)
	fmt.Println(sum)
}

func twoSum(nums []int, target int) []int {
	mm := make(map[int]int)

	for i, v := range nums {
		if x, ok := mm[target-v]; ok {
			return []int{x, i}
		} else {
			mm[v] = i
		}
	}

	return nil
}
