package main

import "container/list"

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	l := list.New()
	l.PushFront(root)

	for l.Len() > 0 {
		len := l.Len()
		for i := 0; i < len; i++ {
			node := l.Remove(l.Back()).(*Node)
			if i < len-1 {
				node.Next = l.Back().Value.(*Node)
			}
			if node.Left != nil {
				l.PushFront(node.Left)
			}
			if node.Right != nil {
				l.PushFront(node.Right)
			}
		}
	}
	return root
}
