package main

import "fmt"

func main() {
	ints := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(ints)
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	mp := make(map[int]bool)
	for _, v := range nums1 {
		mp[v] = true
	}

	var res []int
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			res = append(res, v)
			delete(mp, v)
		}
	}

	return res
}
