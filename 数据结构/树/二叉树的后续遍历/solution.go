package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//        1
//       / \
//      2   3
//      2-->3-->1
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}
