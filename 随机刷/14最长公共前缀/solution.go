package leetcode

import "strings"

//输入：strs = ["flower","flow","flight"]
//输出："fl"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	r := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], r) {
			r = r[:len(r)-1]
			if len(r) == 0 {
				return ""
			}
		}
	}

	return r
}