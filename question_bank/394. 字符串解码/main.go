package main

import (
	"container/list"
	"fmt"
	"strconv"
)

const (
	left  = '['
	right = ']'
)

func main() {
	s := decodeString("3[a]2[bc]")
	fmt.Println(s)
}
func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}
	l := list.New()

	var str string
	for _, v := range s {
		if v != right {
			l.PushFront(string(v))
			continue
		}
		var lp = true
		for lp {
			remove := l.Remove(l.Front()).(string)
			if remove == string(left) {
				var rangeStr = ""

				s2 := l.Remove(l.Front()).(string)
				for _, err := strconv.Atoi(s2); err == nil; {
					rangeStr = s2 + rangeStr
					if l.Len() == 0 {
						break
					} else if _, err = strconv.Atoi(l.Front().Value.(string)); err != nil {
						break
					}

					s2 = l.Remove(l.Front()).(string)
				}
				rangeNum, _ := strconv.Atoi(rangeStr)
				strTmp := str
				for i := 0; i < rangeNum; i++ {
					for _, stmp := range strTmp {
						l.PushFront(string(stmp))
					}
				}
				str = ""
				lp = false
				continue
			}
			str = remove + str
		}
	}
	var tmpStr string
	for e := l.Back(); e != nil; e = e.Prev() {
		tmpStr += e.Value.(string)
	}
	return tmpStr
}
