package leetcode

import (
	"sort"
)
//动态规划问题
// dp[i]代表兑换面额为i的最小硬币数
// dp[i] = dp[i-coin] + 1
// 2
// [1,2,4]
// 默认dp[i]=-1 也就是没有能兑换的

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i:=1; i < len(dp); i++ {
		dp[i] = -1
	}
	for i:=1; i<=amount; i++ {
		for k := 0; k < len(coins); k++ {
			temp := 0
			if i-coins[k] >=0 {
				if dp[i-coins[k]] == -1 {
					continue
				}
				temp = dp[i-coins[k]] + 1
			} else {
				break
			}
			if temp < dp[i] {
				dp[i] = temp
			} else if dp[i] == -1 && temp > 0{
				dp[i] = temp
			}
		}
	}
	return dp[amount]
}
