package leetcode

func maxProfit(prices []int) int {
	if len(prices) <=1 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	for i:=1;i<len(prices);i++{
		if prices[i]>prices[i-1] {
			dp[i] = dp[i-1]+ (prices[i]-prices[i-1])
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

