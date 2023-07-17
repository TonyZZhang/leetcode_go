package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodeLayer := make([][]*TreeNode, 0)
	nodeLayer = append(nodeLayer, []*TreeNode{root})
	res := 1

	for len(nodeLayer) != 0 {
		lastLayer := make([]*TreeNode, 0)
		for i:=0; i<len(nodeLayer[len(nodeLayer)-1]); i++ {
			if nodeLayer[len(nodeLayer)-1][i] == nil {
				continue
			}
			if nodeLayer[len(nodeLayer)-1][i].Left == nil &&
				nodeLayer[len(nodeLayer)-1][i].Right == nil {
				continue
			}
			lastLayer=append(lastLayer, nodeLayer[len(nodeLayer)-1][i].Left)
			lastLayer=append(lastLayer, nodeLayer[len(nodeLayer)-1][i].Right)
		}
		nodeLayer = nodeLayer[:len(nodeLayer)-1]
		if len(lastLayer) != 0 {
			nodeLayer = append(nodeLayer, lastLayer)
			res = big(res, len(lastLayer))
		}
	}
	return res
}

func big(a,b int) int {
	if a > b {
		return a
	}
	return b
}

