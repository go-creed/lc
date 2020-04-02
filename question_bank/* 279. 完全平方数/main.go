package main

import (
	"fmt"
	"math"
)

//示例 1:
//
//输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//示例 2:
//
//输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.
//

func main() {

	squares := numSquares(12)
	fmt.Println(squares)

}

func numSquares(n int) int {

	sqrt := int(math.Sqrt(float64(12)))

	for i := 1; i <= sqrt; i++ {

	}
	fmt.Println(sqrt)

	return 0
}
