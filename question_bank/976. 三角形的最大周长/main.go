package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	perimeter := largestPerimeter([]int{3, 6, 2, 3})
	fmt.Println(perimeter)
}

func add(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

//三角形 任意两边相加小于第三边
func largestPerimeter(A []int) int {
	if len(A) < 3 {
		return 0
	}
	//排序A
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	res := 0
	for i := 0; i+2 < len(A); i++ {
		a := A[i]
		b := A[i+1]
		c := A[i+2]
		if add(a, b, c) {
			res = int(math.Max(float64(res), float64(a+b+c)))
		}
	}
	return res
}
