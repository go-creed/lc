package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := findPairs([]int{1, 1, 1, 1, 1}, 0)
	fmt.Println(pairs)
}

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)
	mp := make(map[int]bool)
	res := 0
	for i, v := range nums[:l-1] {
		if i < l && nums[i+1] == v {
			if _, ok := mp[v]; nums[i+1]-v == k && !ok {
				mp[v] = true
				res++
			}
			continue
		}
		right := i + 1
		for right < l {
			if nums[right]-v == k {
				mp[v] = true
				res++
				break
			} else if nums[right]-v > k {
				break
			}
			right++
		}
	}
	return res
}
