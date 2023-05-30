package leetcode

func lengthOfLongestSubstring(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	if len(s) == 1 {
		return 1
	}
	//abcabcbb
	left := 0
	right := 1
	for right < len(s) && left < right{
		for right < len(s) {
			if !contain(s[right], s[left:right]) {
				right++
			} else {
				break
			}
		}
		if right-left > res {
			res = right-left
		}
		left++
		if left == right {
			right++
		}
	}
	return res
}

func contain(uint82 uint8, string2 string) bool {
	for k, _ := range string2 {
		if uint82 == string2[k] {
			return true
		}
	}
	return false
}
