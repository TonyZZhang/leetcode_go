package leetcode

import "math"

func findPeakElement(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	for i:= 0 ; i <= len(nums);i++ {
		left := 0
		right := 0
		if i == 0 {
			left = 0-math.MaxInt
		} else {
			left = nums[i-1]
		}

		if i == len(nums) {
			right = 0-math.MaxInt
		} else {
			right = nums[i+1]
		}
		if nums[i]>left && nums[i]>right {
			return i
		}
	}
	return -1
}