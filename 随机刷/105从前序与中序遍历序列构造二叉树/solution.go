package leetcode

//给定两个整数数组 preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := new(TreeNode)
	if len(preorder) == 0 {
		return nil
	}
	root.Val = preorder[0]
	left := len(inorder)
	for i := 0; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			left = i
			break
		}
	}

	newInorderLeft := inorder[:left]
	newInorderRight := make([]int, 0)
	if left == len(inorder)-1 {
		newInorderRight = []int{}
	} else {
		newInorderRight = inorder[left+1:]
	}

	if left == 0 {
		root.Left = nil
	} else {
		newPreorderLeftLen := len(newInorderLeft)
		newPreorderLeft := preorder[1:1+newPreorderLeftLen]
		root.Left = buildTree(newPreorderLeft, newInorderLeft)
	}

	if len(newInorderRight) == 0 {
		root.Right = nil
	} else {
		newPreorderLeftLen := len(newInorderLeft)
		newPreorderRight := preorder[1+newPreorderLeftLen:]
		root.Right = buildTree(newPreorderRight, newInorderRight)
	}

	return root
}

