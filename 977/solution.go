package leetcode

func sortedSquares(nums []int) []int {
	var result []int
	result = make([]int, len(nums))
	right := len(nums) - 1
	left := 0
	for i := len(nums)-1; i >= 0; i-- {
		if nums[right] * nums[right] > nums[left] * nums[left] {
			result[i] = nums[right] * nums[right]
			right = right - 1
		} else {
			result[i] = nums[left] * nums[left]
			left = left + 1
		}
	}
	return result
}

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
