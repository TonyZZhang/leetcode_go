package leetcode

import "fmt"
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var r [][]int
	var res [][]int
	var path []int
	dfs(root, path, &res)
	for i:= 0; i<len(res); i++{
		var temp int
		for k:=0; k<len(res[i]); k++ {
			temp += res[i][k]
		}
		if temp == targetSum {
			r = append(r, res[i])
		}
	}
	return r
}

func dfs(root *TreeNode, path[]int, res *[][]int){
	if root == nil {
		return
	}
	newPath := make([]int, len(path)+1)
	copy(newPath, path)
	newPath[len(path)] = root.Val
	if root.Left == nil && root.Right == nil {
		*res = append(*res, newPath)
	} else {
		dfs(root.Left, newPath, res)
		dfs(root.Right, newPath, res)
	}
}