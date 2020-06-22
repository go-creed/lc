package main

import (
	"fmt"

	. "lc/common/tri"
)

func main() {
	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	fmt.Println(tree)
}

func buildPos(inorder []int, val int) int {
	var pos int
	for i, v := range inorder {
		if v == val {
			pos = i
			break
		}
	}
	return pos
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var num = 0
	return buildTreeHelper(preorder, inorder, &num)
}

func buildTreeHelper(preorder []int, inorder []int, num *int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	val := preorder[*num]
	root := &TreeNode{
		Val: val,
	}

	pos := buildPos(inorder, val)

	if pos >= 1 {
		*num++
		root.Left = buildTreeHelper(preorder, inorder[:pos], num)
	}
	if pos+1 < len(inorder) {
		*num++
		root.Right = buildTreeHelper(preorder, inorder[pos+1:], num)
	}
	return root
}
