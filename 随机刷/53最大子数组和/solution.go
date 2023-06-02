package leetcode

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
func maxSubArray(nums []int) int {
	var res, temp int
	count := 0
	neRes := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			count++
			if neRes < nums[i] {
				neRes = nums[i]
			}
		}
		if temp + nums[i] < 0 {
			temp = 0
			continue
		} else {
			temp = temp + nums[i]
			if temp > res {
				res = temp
			}
		}
	}

	if count == len(nums) {
		return neRes
	}
	return res
}
