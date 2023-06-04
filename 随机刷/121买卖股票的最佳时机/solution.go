package leetcode
//[2,1,4]
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var max, res int
	max = prices[len(prices)-1]
	for i := len(prices)-2; i >= 0; i-- {
		if prices[i] < max {
			temp := max - prices[i]
			if temp > res {
				res = max - prices[i]
			}
		} else {
			max = prices[i]
		}
	}
	return res
}
