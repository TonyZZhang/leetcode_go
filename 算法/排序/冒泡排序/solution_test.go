package leetcode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var in []int
	in = append(in, 1)
	in = append(in, 5)
	in = append(in, 2)
	in = append(in, 4)
	in = append(in, 67)
	in = append(in, 3)
	fmt.Println(quickSort(in))
}

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

func selectSort(x []int) []int {
	for i := 0; i < len(x); i++ {
		min := x[i]
		needSwap := i
		for j := i; j < len(x); j++ {
			if x[j] < min {
				min = x[j]
				needSwap = j
			}
		}
		minX := x[needSwap]
		x[needSwap] = x[i]
		x[i] = minX

	}
	return x
}
