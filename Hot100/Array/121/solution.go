package leetcode

dp[i] = dp[i-1] + prices[i]-prices[i-1]

func maxProfit(prices []int) int {
	var res int
	dp := make([]int, len(prices)+1)
	dp[0] = 0
	for i:=1; i<len(prices);i++ {
		if prices[i]-prices[i-1] > 0 {
			dp[i] = dp[i-1]+prices[i]-prices[i-1]
			res = max(res, dp[i])
		}else {
			dp[i] = dp[i-1]
			res = max(res, dp[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}