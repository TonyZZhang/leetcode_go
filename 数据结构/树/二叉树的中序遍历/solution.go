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
//
//       1
//      / \
//     2   3
//     2-->1-->3
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func inorderTraversal_(root *TreeNode) []int {
	var res []int
	if root != nil && root.Left != nil {
		left := inorderTraversal(root.Left)
		res = append(res, left...)
	}
	if root != nil {
		res = append(res, root.Val)
	}
	if root != nil && root.Right != nil {
		right := inorderTraversal(root.Right)
		res = append(res, right...)
	}
	return res
}
