package main

import (
	"fmt"
	"math"
)

func main() {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(profit)
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[1][i] = -prices[i]
			continue
		}
		dp[0][i] = int(math.Max(float64(dp[0][i-1]), float64(dp[1][i-1]+prices[i])))
		dp[1][i] = int(math.Max(float64(dp[1][i-1]), float64(-prices[i])))
	}
	return dp[0][n-1]
}
