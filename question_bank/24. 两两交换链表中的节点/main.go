package main

//24. 两两交换链表中的节点
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.
//通过次数86,811提交次数133,631
import (
	"fmt"
	. "lc/common/lik"
)

func main() {
	head := Generate("[1,2,3]")
	pairs := swapPairs(head)
	fmt.Println(pairs)
}

//如何交换两个节点
//假设这里是2个节点
//1-2
//2-1
//head->1->2->3->4
//head->2->1->3->4
func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{ //添加一个哨兵节点
		Next: head,
	}

	pre := head
	cur := head.Next

	for ; pre.Next != nil; {
		next := cur.Next
		if next == nil {
			break
		}
		nnext := cur.Next.Next
		cur.Next.Next = nil
		cur.Next = nil
		pre.Next = next
		pre.Next.Next = cur
		pre.Next.Next.Next = nnext
		//首轮交换完毕 此时pre节点应该前进
		pre = cur
		cur = nnext
	}

	return head.Next
}
