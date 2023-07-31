package leetcode
//已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
//若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
//若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
//给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
//
//你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//二分查找的变种
func findMin(nums []int) int {
	lens := len(nums)
	middle := lens/2
	//if lens=3 middle=1, lens=4 middle=2
	left := nums[:middle]
	right := nums[middle:]
	var res int
	if len(left) == 0 {
		return right[0]
	}
	if len(right) == 0 {
		return left[0]
	}
	if left[0] == left[len(left)-1] {
		return min(left[0], findMin(right))
	} else if right[0] == right[len(left)-1]  {
		return min(findMin(left), right[0])
	} else if left[0] < left[len(left)-1] {
		res = min(left[0], findMin(right))
	} else if right[0] < right[len(left)-1] {
		res = min(findMin(left), right[0])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

