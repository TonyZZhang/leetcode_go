package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func sumNumbers(root *TreeNode) int {
	res := nodes(root)
	var nums []int
	for i := 0; i < len(res); i++ {
		var temp int
		for k := len(res[i])-1; k >= 0; k-- {
			if k != len(res[i])-1 {
				temp = res[i][k]*10 + temp
			} else {
				temp = res[i][k]
			}
		}
		nums = append(nums, temp)
	}
	r := 0
	for i := 0; i < len(nums); i++ {
		r += nums[i]
	}
	return r
}

//二叉树的所有路径
//func binaryTreePaths(root *TreeNode) []string {
//	var result []string
//	if root == nil {
//		return result
//	}
//	dfs(root, "", &result)
//	return result
//}
//
//func dfs(root *TreeNode, path string, result *[]string) {
//	if root.Left == nil && root.Right == nil {
//		*result = append(*result, path+strconv.Itoa(root.Val))
//		return
//	}
//	if root.Left != nil {
//		dfs(root.Left, path+strconv.Itoa(root.Val)+"->", result)
//	}
//	if root.Right != nil {
//		dfs(root.Right, path+strconv.Itoa(root.Val)+"->", result)
//	}
//}

func nodes(root *TreeNode)[][]int{
	var res [][]int
	var path []int
	if root == nil {
		return nil
	}
	dfsv2(root, path, &res)
	return res
}

func dfsv2(root *TreeNode, path []int, res *[][]int) {
	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		*res = append(*res, path)
		return
	}

	if root.Left != nil {
		path = append(path, root.Left.Val)
		dfsv2(root.Left, path, res)
	}

	if root.Right != nil {
		path = append(path, root.Right.Val)
		dfsv2(root.Right, path, res)
	}
}

