package leetcode

func moveZeroes(nums []int)  {
	zeroNums := 0
	for _, v := range nums {
		if v == 0 {
			zeroNums++
		}
	}

	for i:=0; i<zeroNums; i++ {
		for k:=0; k<len(nums)-1; k++ {
			if nums[k] == 0 {
				nums[k]= nums[k+1]
				nums[k+1] = 0
			}
		}
	}
}
