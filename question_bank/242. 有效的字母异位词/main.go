package main

func main() {
	
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := make([]int, 26, 26)
	for _, v := range s {
		table[v-'a'] ++
	}
	for _, v := range t {
		table[v-'a'] --
		if table[v-'a'] < 0 {
			return false
		}
	}
	return true
}