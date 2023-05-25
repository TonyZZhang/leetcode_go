package leetcode

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var res int
	res = min(height[left], height[right]) * (right - left)
	for left < right {
		if height[left] < height[right] {
			left++
			curRes := min(height[left], height[right]) * (right - left)
			if curRes > res {
				res = curRes
			}
		} else {
			right--
			curRes := min(height[left], height[right]) * (right - left)
			if curRes > res {
				res = curRes
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}