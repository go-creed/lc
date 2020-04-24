package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	binary := addBinary("1010", "1011")
	fmt.Println(binary)
}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)

	max := int(math.Max(float64(aLen), float64(bLen)))
	if aLen < max {
		for i := 0; i < max-aLen; i++ {
			a = "0" + a
		}
	}
	if bLen < max {
		for i := 0; i < max-bLen; i++ {
			b = "0" + b
		}
	}

	next := 0
	var res string

	for i := max - 1; i >= 0; i-- {
		var now string
		a_ := a[i]
		b_ := b[i]
		if a_ == '1' && b_ == '1' {
			if next == 0 {
				now = "0"
			} else {
				now = "1"
			}
			next = 1
		} else if a_ == '0' && b_ == '0' {
			if next == 1 {
				now = "1"
			} else {
				now = "0"
			}
			next = 0
		} else {
			if next == 0 {
				now = "1"
				next = 0
			} else {
				now = "0"
				next = 1
			}
		}
		res = now + res
	}
	if next != 0 {
		res = strconv.Itoa(next) + res
	}
	return res
}
