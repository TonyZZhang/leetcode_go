package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//计算二叉树的高度
//            = 0   p是空节点
// height(p)
//            = max(height(p.left),height(p.right)) + 1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(heightTree(root.Left)-heightTree(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func heightTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(heightTree(root.Left), heightTree(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
