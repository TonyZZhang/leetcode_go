package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	swapChildren(root)
	for len(stack) > 0 {
		currentLen := len(stack)
		for _, v := range stack {
			if v.Left != nil {
				swapChildren(v.Left)
				stack = append(stack, v.Left)
			}
			if v.Right != nil {
				swapChildren(v.Right)
				stack = append(stack, v.Right)
			}
		}
		stack = stack[currentLen:]
	}
	return root
}

func swapChildren(root *TreeNode) {
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
}
