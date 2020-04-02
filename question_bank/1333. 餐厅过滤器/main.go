package main

import (
	"fmt"
	"sort"
)

func main() {
	restaurants := [][]int{{1, 4, 1, 40, 10}, {2, 8, 0, 50, 5}, {3, 8, 1, 30, 4}, {4, 10, 0, 10, 3}, {5, 1, 1, 15, 1}}

	veganFriendly := 0
	maxPrice := 50
	maxDistance := 10
	ints := filterRestaurants(restaurants, veganFriendly, maxPrice, maxDistance)

	fmt.Println(ints)
}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	sort.Slice(restaurants, func(i, j int) bool {
		if restaurants[i][1]==restaurants[j][1]{
			return restaurants[i][0]>restaurants[j][0]
		}
		return restaurants[i][1] > restaurants[j][1]
	})
	var res []int
	for _, v := range restaurants {
		if veganFriendly == 1 {
			if v[2] == veganFriendly && v[3] <= maxPrice && v[4] <= maxDistance {
				res = append(res, v[0])
			}
		} else {
			if v[3] <= maxPrice && v[4] <= maxDistance {
				res = append(res, v[0])
			}
		}
	}
	return res
}
