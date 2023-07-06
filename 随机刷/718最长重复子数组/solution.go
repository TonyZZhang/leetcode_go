package leetcode

func findLength(nums1 []int, nums2[]int) int {
	dp := make([][]int, len(nums1)+1)
	for i:=0; i<len(dp); i++{
		dp[i] = make([]int, len(nums2)+1)
	}
	r := 0
	for i := 0; i<len(dp); i++ {
		for k:= 0; k<len(dp[i]); k++ {
			if nums1[i-1] == nums2[k-1] {
				dp[i][k] = dp[i-1][k-1]+1
				r = big(r, dp[i][k])
			}
		}
	}
	return r
}

func big(a, b int)int{
	if a > b{
		return a
	}
	return b
}

