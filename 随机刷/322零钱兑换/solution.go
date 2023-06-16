package leetcode

import (
	"sort"
	"math"
)
//动态规划问题
// dp[i]代表兑换面额为i的最大硬币数
// dp[i] = dp[i-coin] + 1
// [1,2,4]

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount)
	dp[0] = 0
	for i:=0; i<=amount; i++ {
		for k := 0; k < len(coins); k++ {
			temp := 0
			if i-coins[k] >=0 {
				temp = dp[i-coins[k]] + 1
			}
			if temp > dp[i] {
				dp[i] = temp
			}
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}
