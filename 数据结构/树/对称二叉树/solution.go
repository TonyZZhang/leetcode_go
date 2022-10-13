package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//解题思路：使用队列来实现
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	var queue []*TreeNode
	if root.Left != nil {
		queue = append(queue, root.Left)
	}
	if root.Right != nil {
		queue = append(queue, root.Right)
	}
	if len(queue) != 2 {
		return false
	}

	for len(queue) > 1 {
		if queue[0].Left != nil && queue[1].Right == nil {
			return false
		}
		if queue[0].Left == nil && queue[1].Right != nil {
			return false
		}
		if queue[0].Left != nil && queue[1].Right != nil {
			queue = append(queue, queue[0].Left)
			queue = append(queue, queue[1].Right)
		}
		if queue[0].Right != nil && queue[1].Left == nil {
			return false
		}
		if queue[0].Right == nil && queue[1].Left != nil {
			return false
		}
		if queue[0].Right != nil && queue[1].Left != nil {
			queue = append(queue, queue[0].Right)
			queue = append(queue, queue[1].Left)
		}
		if queue[0] == nil && queue[1] != nil {
			return false
		}
		if queue[0] != nil && queue[1] == nil {
			return false
		}
		if queue[0].Val != queue[1].Val {
			return false
		}
		queue = queue[2:]
	}
	return true
}
