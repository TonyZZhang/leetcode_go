package leetcode

func minPathSum(grid [][]int) int {
	x := len(grid)-1
	if x < 0 {
		return 0
	}
	y := len(grid[0])-1
	if y < 0 {
		return 0
	}
	return dp(x, y, grid)
}

func dp(x, y int, grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[i]))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(dp); i++ {
		for k := 0; k < len(dp[i]); k++ {
			if i == 0 && k == 0 {
				continue
			}
			if i == 0 && k != 0 {
				dp[i][k] = dp[i][k-1] + grid[i][k]
				continue
			}

			if i != 0 && k == 0 {
				dp[i][k] = dp[i-1][k] + grid[i][k]
				continue
			}
			dp[i][k] = min(dp[i-1][k], dp[i][k-1]) + grid[i][k]
		}
	}
	return dp[x][y]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}