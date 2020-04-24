package main

import "fmt"

func main() {
	isomorphic := isIsomorphic("aa", "aa")&&isIsomorphic("aa","aa")
	fmt.Println(isomorphic)
}


func isIsomorphic(s string, t string) bool {
	l := len(s)
	var hashSet = make(map[string]string)
	var res bool
	for i := 0; i < l; i++ {
		ss := fmt.Sprintf("%c", s[i])
		tt := fmt.Sprintf("%c", t[i])
		var v string
		var ok bool
		if v, ok = hashSet[ss]; !ok {
			hashSet[ss] = tt
			v = tt
		}
		if v != tt {
			return false
		}
	}

	return !res
}
