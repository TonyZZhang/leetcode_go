package leetcode

//3 1 4
//1 3 4
//主要在于边缘处理上
//从右边找到第一个比base小的，从左边找到第一个比base大的，交换
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var left, right int
	left = 0
	right = len(nums) - 1
	base := nums[0]
	for left < right {
		for nums[right] >= base && right > left {
			right--
		}
		// right = 1
		for nums[left] <= base && right > left {
			left++
		}
		//left = 0
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp
	}

	nums[0] = nums[left]
	nums[left] = base

	quickSort(nums[0:right])
	quickSort(nums[right+1:])
	return nums
}

//2 1 3
func quickSortV2(nums[]int) []int {
	if len(nums) <= 1 {
		return nums
	}

	flag := nums[0]
	left := 0
	right := len(nums) - 1

	for left < right {
		for nums[right] > flag {
			right--
		}

		for nums[left] < flag {
			left++
		}
		temp := nums[left]
		nums[left] = nums[right]
		nums[right] = temp
	}
	nums[left] = flag
	if left == 0 {
		return nums
	}
	quickSortV2(nums[:left-1])
	quickSortV2(nums[left+1:])
	return nums
}