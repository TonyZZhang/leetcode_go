package leetcode

func uniquePaths(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	if a == 0 || b == 0 {
		return 1
	}
	dp := make([][]int, a)
	for i := 0; i < a; i++ {
		dp[i] = make([]int, b)
	}
	//dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			if i == 0 {
				dp[i][j] = 1
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[a][b]
}

