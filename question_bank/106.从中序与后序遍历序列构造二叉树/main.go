package main

import (
	"fmt"

	. "lc/common/tri"
)

func main() {
	tree := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(tree)
}

func buildPos(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	var num = 0
	return buildTreeHelper(inorder, postorder, &num)
}

func buildTreeHelper(inorder []int, postorder []int, num *int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1-*num]
	root := &TreeNode{Val: val}

	pos := buildPos(inorder, val)
	if pos+1 < len(inorder) {
		*num++
		root.Right = buildTreeHelper(inorder[pos+1:], postorder, num)
	}
	if pos >= 1 {
		*num++
		root.Left = buildTreeHelper(inorder[:pos], postorder, num)
	}

	return root

}
