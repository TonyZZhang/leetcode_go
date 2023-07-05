package leetcode

func matrixSearch(m [][]int, target int) bool {
	x := rowSearch(m, target)
	if len(x) == 0 {
		return false
	}
	return colSearch(x[0], target)
}

func rowSearch(m [][]int, target int)[][]int {
	first := m[0][0]
	middle := m[len(m)/2][0]
	if  first<=target && target<middle {
		m = m[:(len(m))/2]
	} else if middle <= target{
		m = m[(len(m))/2:]
	} else {
		return [][]int{}
	}
	if len(m) == 1 {
		return m
	}
	return rowSearch(m, target)
}

func colSearch(s[]int, target int) bool {
	if target < s[0] || target > s[len(s)-1] {
		return false
	}

	middle := len(s)/2
	if target >= s[0] && target < s[middle] {
		s = s[:middle]
	} else if target >= s[middle] {
		s = s[middle:]
	}
	if len(s) == 1 {
		return true
	}
	return colSearch(s, target)
}

