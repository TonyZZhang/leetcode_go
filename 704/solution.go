package search

func search(nums []int, target int) int {
	var left, right int
	left = 0
	right = len(nums) - 1
	for left <= right {
		middle := left + ((right - left) / 2)
		if nums[middle] < target {
			left = middle + 1
		}else if nums[middle] > target {
			right = middle - 1
		}else {
			return middle
		}
	}
	return -1
}

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
//写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回-1。

//输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
