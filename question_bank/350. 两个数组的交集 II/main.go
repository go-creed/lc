package main

func main() {
	intersect([]int{1, 2, 2, 1}, []int{2, 2})
}

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	set := make(map[int]int, 0)
	for _, v := range nums1 {
		set[v]++
	}

	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			set[v]--
			res = append(res, v)
			if set[v] == 0 {
				delete(set, v)
			}
		}
	}

	return res
}
