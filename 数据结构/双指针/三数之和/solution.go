package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	var i, j int
	j = i + 1
	k := len(nums) - 1
	for i < k {
		a := nums[i]
		if a > 0 {
			return result
		}
		if j == k {
			return result
		}
		b := nums[j]
		c := nums[k]
		if a + b + c == 0 {
			temp := []int{a, b, c}
			result = append(result, temp)
			i++
			j = i+1
			k = len(nums) - 1
			continue
		}
		if a + b + c > 0 {
			k--
		}
		if a + b + c < 0 {
			j++
		}
		if j == k {
			i++
			j = i+1
			k = len(nums) - 1
		}
	}
	return result
}
