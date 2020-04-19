package main

import "fmt"

func main() {
	element := removeElement([]int{0,1,2,2,3,0,4,2}, 2)
	fmt.Println(element)
}

func removeElement(nums []int, val int) int {
	var j int
	for i, num := range nums {
		if num != val { //找到不为3的
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums)
	return j
}
