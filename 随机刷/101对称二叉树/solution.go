package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	return isSymmetric(root.Left) && isSymmetric(root.Right)
}
