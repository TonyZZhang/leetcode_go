package leetcode

//给定 n 节台阶，每次可以走一步或走两步，求一共有多少种方式可以走完这些台阶。
//状态转移方程
//dp[i]=dp[i-1]+dp[i-2]
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
