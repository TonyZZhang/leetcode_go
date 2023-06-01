package leetcode

import "sort"

//待优化
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[k]
}
