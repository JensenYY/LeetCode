package main

/**
 * Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	deep := 0
	if root == nil {
		return deep
	}
	deep++
	lDeep := maxDepth(root.Left)
	rDeep := maxDepth(root.Right)
	if lDeep > rDeep {
		deep += lDeep
	} else {
		deep += rDeep
	}
	return deep
}
