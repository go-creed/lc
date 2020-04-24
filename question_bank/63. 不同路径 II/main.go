package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0,},
		{0, 1, 0,},
		{0, 0, 0,},
	}
	//obstacleGrid =[][]int{
	//	{1},
	//}
	obstacles := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println(obstacles)
}

//常用动态规划解法
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0]==1||obstacleGrid[m-1][n-1]==1{
		return 0
	}

	mv :=-1
	for i := range obstacleGrid[0] {
		if obstacleGrid[0][i]==1{
			mv=0
		}
		obstacleGrid[0][i] = mv
	}
	mv = -1
	for i := range obstacleGrid {
		if obstacleGrid[i][0] ==1{
			mv = 0
		}
		obstacleGrid[i][0] = mv
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i][j-1] + obstacleGrid[i-1][j]
		}
	}

	return -obstacleGrid[m-1][n-1]

}

//todo 滚动更新
