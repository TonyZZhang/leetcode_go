package leetcode

func lotNums(nums []int) int {
	ca := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == ca {
			count++
		} else {
			count--
		}

		if count == 0 {
			ca = nums[i]
		}
	}
	return ca
}
