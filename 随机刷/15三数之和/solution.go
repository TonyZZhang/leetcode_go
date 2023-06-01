package leetcode

import (
	"sort"
	"strconv"
)

func threeSum(nums []int)[][]int{
	res := make([][]int, 0)
	sort.Ints(nums)
	record := make(map[string]int, 0)
	//TODO 去重
	for i := 0; i < len(nums)-2; i++ {
		target := nums[i]
		left := i+1
		right := len(nums)-1
		for left < right {
			if target + nums[left] + nums[right] == 0 {
				key := strconv.Itoa(nums[i])+strconv.Itoa(nums[left])+strconv.Itoa(nums[right])
				_, ok := record[key]
				if ok {
					left++
					continue
				} else {
					record[key] = 0
				}
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				continue
			}
			if target + nums[left] + nums[right] < 0 {
				left++
				continue
			}
			if target + nums[left] + nums[right] > 0 {
				right--
				continue
			}
		}
	}
	return res
}
