package leetcode
//[4,5,6,0,0,0]
//3
//[1,2,3]
//3

//[4,5,1,4,5,6]
func merge(nums1 []int, m int, nums2 []int, n int)  {
	numsRight := len(nums1) - 1
	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[numsRight] = nums2[n-1]
			n--
		} else {
			nums1[numsRight] = nums1[m-1]
			m--
		}
		numsRight--
	}

	for n > 0  {
		nums1[numsRight] = nums2[n-1]
		numsRight--
		n--
	}
}
