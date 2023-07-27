package leetcode

//dp[i] = dp[i-1] + prices[i]-prices[i-1]

func maxProfit(prices []int) int {
	var res int
	minP := prices[0]
	for i := 0; i < len(prices); i++ {
		minP = min(prices[i], minP)
		res = max(res, prices[i]-minP)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}