package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := multiply("12", "8")
	fmt.Println(s)
}

func addZero(str string, n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		builder.WriteString("0")
	}
	builder.WriteString(str)
	return builder.String()
}

func multiply(num1 string, num2 string) string {
	builder := strings.Builder{}
	l1 := len(num1)
	l2 := len(num2)
	if l1 < l2 {
		num1 = addZero(num1, l2-l1)
	} else if l1 > l2 {
		num2 = addZero(num2, l1-l2)
	}

	l := len(num1)
	var next int
	for j := l - 1; j >= 0; j-- {
		n1, _ := strconv.Atoi(string(num1[j]))
		for i := l - 1; i >= 0; i-- {
			n2, _ := strconv.Atoi(string(num2[i]))
			total := n1*n2 + next
			now := total % 10
			next = total / 10
			builder.R
			builder.WriteString(strconv.Itoa(now))
		}
	}
	s := builder.String()
	fmt.Println(s)
	return builder.String()
}
