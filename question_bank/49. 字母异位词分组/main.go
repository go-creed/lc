package main

import "sort"

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	ans := make(map[string][]string)
	for _, str := range strs {
		x := []byte(str)
		sort.Slice(x, func(i, j int) bool {
			return x[i] > x[j]
		})
		key := string(x)
		if _, ok := ans[key]; !ok {
			ans[key] = []string{}
		}
		ans[key] = append(ans[key], str)
	}
	for _, strings := range ans {
		res = append(res, strings)
	}
	return res
}
