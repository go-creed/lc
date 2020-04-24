package main

import . "lc/common/lik"

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	a.Next = b
	b.Next = b
	detectCycle(a)
}
func circle(head *ListNode) *ListNode {
	slow := head
	quick := head

	for ; quick != nil && quick.Next != nil; {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			return slow
		}
	}
	return nil
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	quick := circle(head)
	if quick == nil {
		return nil
	}
	slow := head
	for ; quick != slow; {
		quick = quick.Next
		slow = slow.Next
	}
	return quick
}
