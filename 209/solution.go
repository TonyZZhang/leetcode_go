package leetcode

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 1
		}else {
			return 0
		}
	}

	if nums[0] > target {
		return 1
	}
	left := 0
	right := 1
	result := len(nums) + 1

	for right < len(nums) {
		temp := 0
		for i := left; i <= right ; i++ {
			temp = temp + nums[i]
		}
		if temp < target {
			right++
		}else {
			if right-left+1 < result {
				result = right-left+1
			}
			left++
		}
	}

	if result == len(nums) + 1 {
		return 0
	}

	return result
}

//给定一个含n个正整数的数组和一个正整数 target 。
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组

