package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	// 将整数数组转换为字符串数组
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}

	// 根据题目要求对字符串数组进行排序
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// 如果数组中的元素全是 0，则直接返回 "0"
	if strs[0] == "0" {
		return "0"
	}

	// 将排序后的字符串数组拼接成一个字符串并返回
	return strings.Join(strs, "")
}
