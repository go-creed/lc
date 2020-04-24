package main

import (
	"fmt"
	. "lc/common/lik"
)

func main() {
	head := Generate("[1,2,3,4,5]")
	end := removeNthFromEnd(head, 5)
	fmt.Println(end)
}

//要删除倒数的第n个节点，可以用双指针，前指针先走n个节点，然后同时走，直到n在尾部
//考虑到删除第一个节点可能会遇到问题，这里添加了一个头结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = &ListNode{Next: head}
	pre := head
	for cur := head; cur.Next != nil; cur = cur.Next {
		if n > 0 {
			n--
			continue
		}
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return head.Next
}
