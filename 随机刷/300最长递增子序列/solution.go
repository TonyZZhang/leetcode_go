package leetcode
//动态规划
//dp[i] = dp[i-1]+1 if dp[i]>dp[i-1]
func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)
	var res int
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				temp := dp[j]+1
				if temp > dp[i] {
					dp[i] = temp
				}
			}
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
