package 在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (right+left) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	rightNum := right

	//5,7,7,8,8,10
	left = 0
	right = len(nums) - 1
	for left < right {
		mid := (right+left+1) / 2
		if nums[mid] == target {
			left = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	leftNum := left
	if leftNum > right {
		return []int{-1, -1}
	}
	result := []int{rightNum, leftNum}
	return result
}

