package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//使用队列实现
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		currentLen := len(queue)
		for i := 0; i < currentLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[currentLen:]
		result++
	}
	return result
}
