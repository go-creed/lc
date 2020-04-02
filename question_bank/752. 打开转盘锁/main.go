package main

import (
	"container/list"
	"fmt"
)

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	deadends = []string{"8888"}
	target = "0009"
	lock := openLock(deadends, target)
	fmt.Println(lock)
}

func openLock(deadends []string, target string) int {
	initLock := []byte{'0', '0', '0', '0'}
	deadMap := make(map[string]bool)
	for _, v := range deadends {
		deadMap[v] = true
	}
	mp := make(map[string]bool)
	mp["0000"] = true
	used := openLockUsed(deadMap, target, mp, string(initLock))
	return used
}

func openLockUsed(deadends map[string]bool, target string, mp map[string]bool, initLock string) int {
	link := list.New()
	link.PushBack(initLock)
	var res int
	for link.Len() > 0 {
		l := link.Len()
		for i := 0; i < l; i++ {
			str := link.Remove(link.Front()).(string)
			if _, ok := deadends[str]; ok {
				continue
			}
			if str == target {
				return res
			}
			for j := 0; j < 4; j++ {
				for k := 0; k < 2; k++ {
					var tmp string
					if k == 0 {
						tmp = up(j, str)
					} else {
						tmp = down(j, str)
					}
					if _, ok := mp[tmp]; !ok {
						link.PushBack(tmp)
						fmt.Println(tmp)
						mp[tmp] = true
					}
				}
			}
		}
		res++
	}
	return -1
}

func up(pos int, str string) string {
	temp := []byte(str)
	if temp[pos] == '9' {
		temp[pos] = '0'
	} else {
		temp[pos]++
	}
	return string(temp)
}

func down(pos int, str string) string {
	temp := []byte(str)
	if temp[pos] == '0' {
		temp[pos] = '9'
	} else {
		temp[pos]--
	}
	return string(temp)
}
