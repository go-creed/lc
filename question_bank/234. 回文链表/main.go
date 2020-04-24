package main

import (
	"fmt"
	. "lc/common/lik"
)

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	palindrome := isPalindrome(node)
	fmt.Println(palindrome)
}

func isPalindrome(head *ListNode) bool {
	var res []int

	for h := head; h != nil; h = h.Next {
		res = append(res, h.Val)
	}
	right := len(res) - 1 //[0...3]
	left := 0

	for left < right {
		if res[left] != res[right] {
			return false
		}
		left++
		right--
	}

	return true

}

func length(head *ListNode) int {
	var len int

	for h := head; h != nil; h = h.Next {
		len++
	}
	return len
}
