package leetcode


 type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
 }

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}

func helper(node * TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}

