package leetcode

//1,3
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] > target {
			return 0
		} else {
			return 1
		}
	}
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
