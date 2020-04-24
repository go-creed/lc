package main

import "fmt"

func main() {
	duplicate := containsDuplicate([]int{1, 2, 3, 4})

	fmt.Println(duplicate)
}

func containsDuplicate(nums []int) bool {
	var res bool
	var hashSet = make(map[int]int)
	for _, num := range nums {
		if _, ok := hashSet[num]; !ok {
			hashSet[num] = 1
		} else {
			return true
		}
	}
	return res
}
