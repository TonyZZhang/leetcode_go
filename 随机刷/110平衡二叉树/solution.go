package leetcode

import "math"

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(depth(root.Left))-float64(depth(root.Right))) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right))+1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
