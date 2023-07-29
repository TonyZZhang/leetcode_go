package leetcode

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//[-2, -3, 4, -3, 2, 3]

//因为是连续的（包含nums[i]以及前面的）所以直接可以知道dp[i] = dp[i-1]+nums[i]，这里有一个很有趣的地方
//
//dp[i-1]	nums[i]	           要不要加上
//正	     正	           都是正的加上肯定变大啊，跟前面的子序列连上变成更长的子序列
//正	     负	           前面是正的，赶快加上前面的连续子序列，让自己变的大一点
//负	     负	           本来就是负的了，再加上dp[i-1]的负数就更小了，不加，独立门户，自己就是连续子序列
//负	     正	           自己是正的，为什么要跟前面的负数连在一起，独立门户，自己就是连续子序列

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	var res int
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1]+nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
