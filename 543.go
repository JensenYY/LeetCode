package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxElem(first int, args ...int) int {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}

func treeRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxElem(treeRecursive(root.Left), treeRecursive(root.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := treeRecursive(root.Left) + treeRecursive(root.Right)
	return maxElem(ret, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}
