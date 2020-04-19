package main

import "fmt"

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。
//
//
//
//说明:
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

func main() {
	merge([]int{2,0}, 1, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	n1pos := m - 1
	n2pos := n - 1
	for i := len(nums1) - 1; i >= 0 && n2pos >= 0; i-- {
		if n1pos >= 0 && nums1[n1pos] >= nums2[n2pos] {
			nums1[i] = nums1[n1pos]
			n1pos--
		} else {
			nums1[i] = nums2[n2pos]
			n2pos--
		}
	}
	fmt.Println(nums1)
}
