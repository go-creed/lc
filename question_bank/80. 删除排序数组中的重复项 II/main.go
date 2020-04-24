package main

func main() {
	removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
}

const max = 2

func removeDuplicates(nums []int) int {
	var i int = 0;
	for _, v := range nums {
		if i < 2 || v > nums[i-2] {
			nums[i] = v
			i++
		}
	}
	return i
}
