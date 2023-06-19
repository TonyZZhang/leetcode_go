package leetcode

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	dp := make([][]int, 0)
	s := nums[0]
	dp = append(dp, []int{})
	dp = append(dp, []int{s})

	for i := 1; i < len(nums); i++ {
		curLen := len(dp)
		for k := 0; k < curLen; k++ {
			item := dp[k]
			newItem := make([]int,0)
			for v := 0; v < len(item); v++ {
				newItem = append(newItem, item[v])
			}
			newItem = append(newItem, nums[i])
			dp = append(dp, newItem)
		}
	}
	return dp
}

