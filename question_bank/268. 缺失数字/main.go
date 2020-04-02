package main

import "fmt"

func main() {
	number := missingNumber([]int{3, 0, 1})
	fmt.Println(number)
}

func missingNumber(nums []int) int {
	n := len(nums) + 1

	array := make(map[int]bool, n)

	for _, v := range nums {
		array[v] = true
	}

	for i := 0; i < n; i++ {
		if _, ok := array[i]; !ok {
			return i
		}
	}
	return -1
}
