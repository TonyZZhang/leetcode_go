package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isSymmetric(root *TreeNode) bool {
	que := make([]*TreeNode, 0)
	if root == nil {
		return true
	}
	if root.Left != nil {
		que = append(que, root.Left)
	}
	if root.Right != nil {
		que = append(que, root.Right)
	}
	if len(que)%2 != 0 {
		return false
	}

	curQueLen := len(que)
	for i:= 0; i<curQueLen; i++{
		if que[i].Left != nil {
			que = append(que, que[i].Left)
		}
		if que[i].Right != nil {
			que = append(que, que[i].Right)
		}
		if i == curQueLen-1 {
			curQueLen = len(que)
			l := curQueLen-1-i
			if l%2 != 0 {
				return false
			}
			for k:= i+1; k <= i+(l/2); k++ {
				if que[k].Val != que[curQueLen-1-k].Val {
					return false
				}
			}
		}
	}

	return true
}

