package leetcode

import "math"

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	if matrix[0][0] == '1' {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	//dp[0][0] = matrix[0][0]
	r := 0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i-1][j-1] == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			} else {
				//todo:从dp[i][j]找两个边，全部为‘1’才能用下面的，否在为1或者0
				dp[i][j] = int((math.Sqrt(float64(dp[i-1][j-1]))+1)*(math.Sqrt(float64(dp[i-1][j-1]))+1))
				r = big(r, dp[i][j])
			}
		}
	}
	return r
}

func big(a,b int) int {
	if a > b {
		return a
	}
	return b
}
