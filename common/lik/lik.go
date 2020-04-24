package lik

import (
	"log"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Lister interface {
	Generate(str string) *ListNode
}

func Generate(str string) *ListNode {
	array := strToArray(str)
	root := &ListNode{}
	head := root
	for _, v := range array {
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next
	}
	return root.Next
}

func strToArray(str string) []int {
	str = strings.Trim(str, "[")
	str = strings.Trim(str, "]")
	strs := strings.Split(str, ",")
	var res []int
	for _, v := range strs {
		i, e := strconv.Atoi(v)
		if e != nil {
			log.Fatal(e)
		}
		res = append(res, i)
	}
	return res
}
