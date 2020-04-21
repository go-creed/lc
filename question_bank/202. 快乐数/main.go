package main

import "fmt"

func main() {
	happy := isHappy(19)
	fmt.Println(happy)
}

func getNext(n int) int {
	var totalSUm = 0
	for n > 0 {
		d := n % 10
		n = n / 10
		totalSUm += d * d
	}
	return totalSUm
}

func isHappy(n int) bool {
	hashSet := make(map[int]struct{})
	for {
		if _, ok := hashSet[n]; !ok && n != 1 {
			hashSet[n] = struct{}{}
			n = getNext(n)
		}else {
			break
		}
	}
	return n == 1
}
