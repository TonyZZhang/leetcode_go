package 盛水最多的容器

func waterContainer(nums []int) int {
	left := 0
	right := len(nums)-1
	res := min(nums[left], nums[right]) * (right-left)
	for left < right{
		curArea := min(nums[left], nums[right]) * (right-left)
		res = max(res, curArea)
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}