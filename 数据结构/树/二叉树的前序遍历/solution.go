package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//        1
//       / \
//          2
//         / \
//        3
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
//递归写法
func preorderTraversal_(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, root.Val)
	} else {
		return res
	}
	if root.Left != nil {
		left := preorderTraversal(root.Left)
		res = append(res, left...)
	}
	if root.Right != nil {
		right := preorderTraversal(root.Right)
		res = append(res, right...)
	}
	return res
}
