package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	var pos int
	for _, num := range nums {
		if num != 0 {
			nums[pos] = num
			pos++
		}
	}
	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(pos)
	fmt.Println(nums)
}
