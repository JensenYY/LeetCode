package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumUp(root *TreeNode, pre, sum int) int {
	if root == nil {
		return 0
	}
	count := pre + root.Val
	if count == sum {
		return 1 + sumUp(root.Left, count, sum) + sumUp(root.Right, count, sum)
	}
	return sumUp(root.Left, count, sum) + sumUp(root.Right, count, sum)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return sumUp(root, 0, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
