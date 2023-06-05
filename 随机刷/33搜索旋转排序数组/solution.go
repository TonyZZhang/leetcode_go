package leetcode

//整数数组nums 按升序排列，数组中的值互不相同。在传递给函数之前，nums在预先未知的某个下标k(o)上进行了旋转，使数组变为[nums [ k ]，
//nums[ k+l], ..., nums [ n-l] , nums [ o], nums [ 1],..., nums [ k-1]](下标从o 开始计数)。例如，[0,1,2,4,5,6,7]在下标3处经旋转后可能变为[4,5,6,7,0,1,2]。
//给你旋转后的数组nums 和一个整数target ，如果nums中存在这个目标值target ，则返回它的下标，否则返回-1。
//二分查找
// [4,5,6,7,0,1,2] 0

func search(nums []int, target int) int{
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[left] {
			if nums[mid] >= target && nums[mid] < target{
				right = mid-1
			} else {
				left = mid+1
			}
		} else {
			if target > nums[mid] &&  target <= nums[right]{
				left = mid+1
			} else {
				right = mid-1
			}
		}
	}

	return -1
}

