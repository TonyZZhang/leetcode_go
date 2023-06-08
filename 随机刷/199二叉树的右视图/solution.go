package leetcode


type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

func popNode (que []*TreeNode) []*TreeNode {
	if len(que) > 1 {
		que = que[1:]
	} else {
		que = []*TreeNode{}
	}

	return que
}

func rightSideView(root *TreeNode) []int {
	var res []int
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) != 0 {
		length := len(que)
		for i := 0; i < length; i++ {
			node := que[0]
			if node != nil {
				if node.Left != nil {
					que = append(que, node.Left)
				}
				if node.Right != nil {
					que = append(que, node.Right)
				}
				if i == length -1 {
					res = append(res, que[0].Val)
				}
			}
			que = que[1:]
		}
	}
	return res
}