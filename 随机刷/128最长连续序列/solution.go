package leetcode

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, v:= range nums{
		numSet[v] = true
	}

	maxLen := 0
	for num := range numSet {
		if numSet[num-1] == false {
			cur := num
			tempLen := 0
			for numSet[cur] {
				tempLen++
				cur++
			}
			maxLen = max(maxLen, tempLen)
		}
	}

	return maxLen
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}