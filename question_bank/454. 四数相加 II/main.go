package main

import (
	"fmt"
)

func main() {
	count := fourSumCount([]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1})
	fmt.Println(count)
}

func sum(a []int, b []int) map[int]int {
	m := make(map[int]int, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			m[a[i]+b[j]] += 1
		}
	}
	return m
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sumAB := sum(A, B)
	fmt.Println(sumAB)
	var res int
	for _, c := range C {
		for _, d := range D {
			x, _ := sumAB[-c-d]
			res += x
		}
	}
	return res
}
